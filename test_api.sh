#!/bin/bash

# Define API URL
API_URL="http://18.188.209.13:8080"

# Test File Upload
echo "Testing file upload..."
curl -X POST -F "myfile=@/Users/abhyudaisingh/Downloads/Behavioral.txt" "$API_URL/upload"   
curl -X POST -F "myfile=@./Behavioral.txt" "$API_URL/upload"

# Test MD5 Check
echo -e "\nTesting MD5 check..."
curl "$API_URL/checkmd5?filename=Behavioral.txt"
