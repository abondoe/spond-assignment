#!/bin/bash

# type CreateRegistrationRequest struct {
# 	FormId       string     `json:"formId"`
# 	MemberTypeId string     `json:"memberTypeId"`
# 	Name         string     `json:"name"`
# 	Email        string     `json:"email"`
# 	PhoneNumber  string     `json:"phoneNumber"`
# 	BirthDate    *time.Time `json:"birthDate"`
# }
#
# Example form body:
# {
#   "clubId": "britsport",
#   "memberTypes": [
#     {
#       "id": "8FE4113D4E4020E0DCF887803A886981",
#       "name": "Active Member"
#     },
#     {
#       "id": "4237C55C5CC3B4B082CBF2540612778E",
#       "name": "Social Member"
#     }
#   ],
#   "formId": "B171388180BC457D9887AD92B6CCFC86",
#   "title": "Coding camp summer 2025",
#   "registrationOpens": "2024-12-16T00:00:00Z"
# }


# Test script to create a registration using the API
curl -X POST http://localhost:8080/registrations \
  -H "Content-Type: application/json" \
  -d '{
    "memberId": "019D8B2358C17A56A95900AB3A48078E",
    "formId": "B171388180BC457D9887AD92B6CCFC86",
    "memberTypeId": "8FE4113D4E4020E0DCF887803A886981",
    "name": "John Doe",
    "email": "john.doe@example.com",
    "phoneNumber": "12345678",
    "birthDate": "1990-01-01T00:00:00Z"
  }'