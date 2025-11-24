require "./query"
require "json"
require "time"

token = "2Ecn0AhG60-gWm_z1wfkZA2"

from_date = Time.local(2025, 12, 30)
formatted_from_date = from_date.to_s("%Y-%m-%d")

to_date = from_date.shift(years: 3) - 1.day
formatted_to_date = to_date.to_s("%Y-%m-%d")

json_text = %({
  "RentalContract": {
    "BaseDate": "#{formatted_from_date}"
  }
})

payload = JSON.parse(json_text)

endpoint = "rental/contracts/100490?$db=TEM-NO"

query = Query.new(endpoint,token)
response = query.put(payload);

puts "Body: #{response}"

endpoint = "rental/contracts/100490/items/api-operations/101969/instances?$db=TEM-NO"
query = Query.new(endpoint,token)

json_text = %({
  "Parameters": {
    "RentalContractItem": {
      "Quantity": "1",
      "Period": {
        "EffectiveDate": "#{formatted_from_date}",
        "ExpirationDate": "#{formatted_to_date}"
      },
      "RecurringRate": {
        "RecurringRateId": "102526"
      },
      "Product": {
        "ProductId": "122605"
      }
    }
  }
})

payload = JSON.parse(json_text)
response = query.post(payload);

data = JSON.parse(response.body)

puts "Body: #{data}"



