using System.Net;
using System.Net.Sockets;
using System.Text;

namespace HttpServer.Services;

public class HttpServer
{
    private readonly TcpListener _listener;

    public HttpServer(int port)
    {
        _listener = new TcpListener(IPAddress.Parse("127.0.0.1"), port);
    }

    public async Task ListenForRequestsAsync()
    {
        int i = 0;
        _listener.Start();
        while (true)
        {
            try
            {
                Console.WriteLine("Waiting for client {0}", i++);
                var client = await _listener.AcceptTcpClientAsync();
                Console.WriteLine("TcpClient accepted");

                var buffer = new byte[10240];
                var stream = client.GetStream();

                var length = await stream.ReadAsync(buffer, 0, buffer.Length);
                var incomingMessage = Encoding.UTF8.GetString(buffer, 0, length);

                Console.WriteLine("Incoming message:");
                Console.WriteLine(incomingMessage);

                var method = GetMethod(incomingMessage);
                Console.WriteLine("Method type: " + method);

                var httpResponse = GetBody(method);

                Console.WriteLine("Response message:");
                Console.WriteLine(httpResponse);

                await stream.WriteAsync(Encoding.UTF8.GetBytes(httpResponse));
                await stream.FlushAsync();
                stream.Close();
                client.Close();
            }
            catch (Exception ex)
            {
                Console.WriteLine("Error occurred: " + ex.Message);
            }

            await Task.Delay(1000);
        }
    }

    private string GetMethod(string incomingMessage)
    {
        var requestLines = incomingMessage.Split(new[] { Environment.NewLine }, StringSplitOptions.None);
        var requestLine = requestLines[0];
        var requestParts = requestLine.Split(' ');
        var method = requestParts[0];
        return method;
    }

    private string GetBody(string method)
    {
        string httpBody;
        var statusLine = "HTTP/1.0 200 OK";

        if (method.Equals("GET", StringComparison.InvariantCultureIgnoreCase))
        {
            httpBody = LoadHtmlFromFile("home.html");
        }
        else if (method.Equals("POST", StringComparison.InvariantCultureIgnoreCase))
        {
            httpBody = LoadHtmlFromFile("send.html");
        }
        else
        {
            statusLine = "HTTP/1.0 405 Method Not Allowed";
            httpBody = $"<html><h1>Method {method} Not Allowed! {DateTime.Now} </h1></html>";
        }

        var httpResonse = statusLine + Environment.NewLine
                                     + "Content-Length: " + Encoding.UTF8.GetByteCount(httpBody) +
                                     Environment.NewLine
                                     + "Content-Type: text/html" + Environment.NewLine
                                     + Environment.NewLine
                                     + httpBody
                                     + Environment.NewLine + Environment.NewLine;

        return httpResonse;
    }

    static async Task SendPostRequest()
    {
        using (HttpClient client = new HttpClient())
        {
            var content = new StringContent("Sample data", Encoding.UTF8, "application/json");
            HttpResponseMessage response = await client.PostAsync("http://127.0.0.1:13000", content);
            string responseBody = await response.Content.ReadAsStringAsync();
            Console.WriteLine("POST Response:");
            Console.WriteLine(responseBody);
        }
    }

    private string LoadHtmlFromFile(string fileName)
    {
        try
        {
            // Assuming the HTML files are located in the same directory as the executable
            string filePath = Path.Combine(AppDomain.CurrentDomain.BaseDirectory, fileName);
            return File.ReadAllText(filePath);
        }
        catch (Exception ex)
        {
            Console.WriteLine($"Error reading file {fileName}: {ex.Message}");
            return $"<html><h1>Error loading page</h1><p>{ex.Message}</p></html>";
        }
    }
}