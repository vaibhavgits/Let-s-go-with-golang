# API

### GET

Used to retrieve data.
Data is sent via URL query parameters.
Example: http://localhost:8000/get?name=Vaibhav.

### POST

Used to send data to the server.
Data is sent in the request body, typically as JSON.
Example: JSON payload {"text": "Hello, world!"}.

### POSTForm

Similar to POST but data is sent as form-encoded.
Commonly used in HTML forms.
Example: Form data name=Vaibhav.

### Form-encoded


- "Form-encoded" refers to the way data is formatted when it's sent from an HTML form to a server using the POST method. 
- Specifically, the data is encoded in a key-value pair format, where each key-value pair corresponds to a form field and its submitted value. This encoding format is known as application/x-www-form-urlencoded.

#### Characteristics of Form-encoded Data
Key-Value Pairs: Each form field is represented as a key-value pair.
Ampersand (&) Separation: Pairs are separated by an ampersand (&).
Equals (=) Assignment: Keys and values are separated by an equals sign (=).
URL Encoding: Special characters are URL-encoded to ensure the data is correctly transmitted (e.g., spaces become %20)


### POST vs POSTForm

1. Data Format:  
   
- POST: Can use various formats like JSON, XML, or form-encoded.
- POSTForm: Specifically uses form-encoded data (application/x-www-form-urlencoded).


2. Typical Usage:
   
- POST: Often used for APIs where JSON or XML is preferred for complex data structures.
- POSTForm: Commonly used for web forms where simple key-value pairs are sent.
  

3. Content-Type Header"
   
- POST: The Content-Type can be set to any appropriate type (e.g., application/json).
- POSTForm: The Content-Type is application/x-www-form-urlencoded.