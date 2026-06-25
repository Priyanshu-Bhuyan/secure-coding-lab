## Environment Variables

The application requires the following environment variables before execution:

```bash
APP_PASSWORD=your_application_password
JWT_SECRET=your_jwt_secret_key
DB_CONN=user:password@tcp(localhost:3306)/dbname
```

If any required secret is not configured, the application exits during startup to prevent insecure execution.
