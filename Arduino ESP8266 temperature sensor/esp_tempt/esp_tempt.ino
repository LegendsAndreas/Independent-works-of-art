// The code for setting up the temperature sensor was aquired from here: https://randomnerdtutorials.com/guide-for-ds18b20-temperature-sensor-with-arduino/
#include <OneWire.h>
#include <DallasTemperature.h>
#include <ESP8266WiFi.h>
#include <WiFiClientSecure.h>
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

  // Connects to MQTT Server
  connectToMQTT();

  // Start up the temperature library
  sensors.begin();
}

void loop(void){ 
  // If WiFi gets diconnected, we try to reestablish a connection
  if (WiFi.status() != WL_CONNECTED) {
    Serial.println("WiFi disconnected, attempting to reconnect...");
    connectToWiFi();
  }

  // call poll() regularly to allow the library to send MQTT keep alive which avoids being disconnected by the broker
  //connectToMQTT();
  mqttClient.poll();

  // Gets the temperature.
  const double temperature = getTemperature();

  // Sends temperature to the MQTT server
  sendTopic(temperature);

  // Disconnects from the MQTT, since there are 29 minutes where it does nothing
  //Serial.println("Disconnecting from the MQTT server, see ya in 29 mins!")

  // PROBLEM (i think): Whenever the disconnect code is not commented out, it does not update the MQTT server.
  // Maybe, it is because it disconnects before the message can be really send. Maybe just adding a delay will work...
  // Disconnects form the WiFi. You cant ONLY disconnect from the WiFi
  //WiFi.disconnect();
  //Serial.println("Disconnecting from the WiFi, see ya in 29 mins!")

  // Add threshold that prints the temperature right away, if it raises by a certain amount

  // Page for changing WiFi, MQTT and such.

  // We wait 29 minutes. The program that extract the information, does it every 30 minutes, so by taking 1 less minute to compute the temperature,
  // allows us to be sure that we always get a new temperature
  delay(MINUTE*30);
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