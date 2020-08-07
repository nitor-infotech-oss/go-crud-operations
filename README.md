# RESTful API Example with golang
This is simple example restful api server only with **gorilla/mux**.  
For simplifying code, this example uses a mock database that is `map[string]interface{}`

## API Endpoint
- http://localhost:8089/api/v1/patient
    - `GET`: get list of patients
    - `POST`: create patient
- http://localhost:8089/api/v1/patient/{name}
    - `GET`: get patient
    - `PUT`: update patient
    - `DELETE`: remove patient

## Data Structure
```json
{
  "name": "dattatray",
  "tel": "7276246800",
  "email": "dattatray.pawar@nitorinfotech.com"
}
```