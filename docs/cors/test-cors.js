// Simple test script to check CORS
fetch("http://localhost:8080/tasks", {
  method: "GET",
  headers: {
    "Content-Type": "application/json",
    Accept: "application/json",
  },
})
  .then((response) => {
    console.log("Success:", response.status);
    return response.json();
  })
  .then((data) => console.log("Data:", data))
  .catch((error) => console.error("CORS Error:", error));
