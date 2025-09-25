# CORS Debugging Checklist

## Changes Made:

1. ✅ Replaced cors library with custom middleware for better debugging
2. ✅ Simplified Angular HTTP client to use default headers
3. ✅ Added comprehensive logging to CORS middleware

## Testing Steps:

### 1. Start the Go server:

```bash
cd /Users/rubenalves/Documents/repos/taskuik/server
go run cmd/server/main.go
```

### 2. Check server logs:

- Look for "CORS Request:" log messages
- Verify the server is running on port 8080

### 3. Test with curl:

```bash
chmod +x test-cors.sh
./test-cors.sh
```

### 4. Start Angular dev server:

```bash
cd /Users/rubenalves/Documents/repos/taskuik/client
ng serve
```

### 5. Open browser and test:

- Go to http://localhost:4200
- Open Developer Tools (F12) → Network tab
- Try to load tasks
- Check for:
  - OPTIONS preflight request
  - GET request to /tasks
  - Response headers in both requests

### 6. Alternative browser test:

- Open cors-test.html in browser
- Click "Test CORS" button
- Check console for detailed error messages

## Common CORS Error Messages and Solutions:

### "Access to fetch at 'http://localhost:8080/tasks' from origin 'http://localhost:4200' has been blocked by CORS policy"

- Server is not responding with proper CORS headers
- Check if server is running and accessible

### "CORS policy: Request header field content-type is not allowed"

- Custom headers require preflight OPTIONS request
- OPTIONS request should respond with Access-Control-Allow-Headers

### "CORS policy: No 'Access-Control-Allow-Origin' header is present"

- Server is not setting the CORS header at all
- Check if CORS middleware is properly applied

### "CORS policy: The request client is not a secure context"

- Usually happens with mixed HTTP/HTTPS content
- Make sure both frontend and backend use same protocol

## Debugging Commands:

### Check if server is responding:

```bash
curl http://localhost:8080/tasks
```

### Check CORS headers:

```bash
curl -I http://localhost:8080/tasks -H "Origin: http://localhost:4200"
```

### Check preflight:

```bash
curl -X OPTIONS http://localhost:8080/tasks -H "Origin: http://localhost:4200" -v
```

## If still having issues:

1. Share the exact error message from browser console
2. Share the Network tab details (request/response headers)
3. Share the server logs showing the CORS requests
