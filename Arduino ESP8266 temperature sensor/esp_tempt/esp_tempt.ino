// The code for setting up the temperature sensor was aquired from here: https://randomnerdtutorials.com/guide-for-ds18b20-temperature-sensor-with-arduino/
#include <OneWire.h>
#include <DallasTemperature.h>
#include <ESP8266WiFi.h>
#include <ArduinoMqttClient.h>
#include "credentials.h"

// 60,000 milliseconds, which equals to a full minute
#define MINUTE 60000

// WiFi part
const char* ssid = SSID;
const char* password = PASSWORD;

// Temperature part
#define ONE_WIRE_BUS 4 // Data wire is conntec to the Arduino ESP pin 4
OneWire oneWire(ONE_WIRE_BUS); // Setup a oneWire instance to communicate with any OneWire devices
DallasTemperature sensors(&oneWire); // Pass our oneWire reference to Dallas Temperature sensor

// MQTT part
WiFiClient wifiClient;
MqttClient mqttClient(wifiClient);

const char* broker = BROKER;
const int   port   = PORT;
const char* topic  = TOPIC;
const char* mqtt_user = MQTT_USERNAME;
const char* mqtt_pass = MQTT_PASSWORD;

void setup(void)
{
  // Initialize serial and wait for port to open
  Serial.begin(115200);
  while (!Serial) {
    ; // Wait for serial port to connect. Needed for native USB port only
  }

  // Connects to WiFi
  connectToWiFi();

  // Start up the temperature library
  sensors.begin();
}

void loop(void){ 
  // If WiFi gets diconnected, we try to reestablish a connection
  if (WiFi.status() != WL_CONNECTED) {
    Serial.println("WiFi disconnected, attempting to reconnect...");
    connectToWiFi();
  }

  // Connects to the MQTT server.
  // After a certain amount of time, you will automatically be disconnect from the MQTT server, so we dont have to actually write a disconnect statement.
  // The "connectToMQTT" function, also checks if we are already connected.
  connectToMQTT();

  // Gets the temperature.
  const double temperature = getTemperature();

  // Sends temperature to the MQTT server
  sendTopic(temperature);

  // Disconnects from the WiFi. We disconnect so that we dont use unnecessary power to stay connected to the WiFi.
  // If we do not have this short delay, the MQTT server will not be updated. I believe this is because we disconnect from the internet, BEFORE, we fully send the message to our MQTT server.
  delay(5000);
  WiFi.disconnect();
  Serial.println("Disconnecting from the WiFi, see ya in 25 mins!");

  // We wait 25 minutes. The program that extract the information, does it every 30 minutes, so by taking 5 less minute to compute the temperature,
  // allows us to be sure that we always get a new temperature
  delay(MINUTE*25);
}

// Establishes a connection to the selected WiFi.
void connectToWiFi(void) {
  // Connecting to WiFi
  Serial.println();
  Serial.print("Connecting to ");
  Serial.println(ssid);
  WiFi.mode(WIFI_STA);
  WiFi.begin(ssid, password);
  while (WiFi.status() != WL_CONNECTED) {
    delay(5000);
    Serial.print(".");
  }

  // Prints IP-address
  Serial.println("");
  Serial.println("WiFi connected");
  Serial.println("IP addresses: "); 
  Serial.println(WiFi.localIP()); 
  Serial.println(WiFi.macAddress());
}

// Uses a temperature sensor to calculate the temperature, with the Dallas temperature library
const double getTemperature(void) {
  // Call sensors.requestTemperatures() to issue a global temperature and Requests to all devices on the bus
  sensors.requestTemperatures(); 
  // Why "byIndex"? You can have more than one IC on the same bus. 0 refers to the first IC on the wire
  const double tempt = sensors.getTempCByIndex(0);
  
  Serial.print("Celsius temperature: ");
  Serial.println(tempt);

  return tempt;
}

// Connects to the MQTT broker
void connectToMQTT(void) {
  Serial.print("Connecting to MQTT broker: ");
  Serial.println(broker);

  mqttClient.setUsernamePassword(mqtt_user, mqtt_pass);
  if (!mqttClient.connect(broker, port)) {
    Serial.print("MQTT connection failed! Error code = ");
    Serial.println(mqttClient.connectError());
    while (1);
  }


}

// Sends a topic to a MQTT broker/server
void sendTopic(const double tempt) {
  mqttClient.beginMessage(topic);
  mqttClient.print(tempt);
  mqttClient.endMessage();
}
