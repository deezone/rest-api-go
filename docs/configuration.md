# rest-api-go

### Configuration

#### Environment Variables
- [Linux: Set Environment Variable](https://www.cyberciti.biz/faq/set-environment-variable-linux/)
```
export REST_API_ENV=production

OR

export REST_API_ENV=development

OR

export REST_API_ENV=test

```

#### Configuration files
Add configuration files to:
- `config/config.development.json`
- `config/config.production.json`

*Sample*
```
{
  "Version": "v0.01",
  "ReleaseDate": "2018-01-06T18:00:00",
  "Port": 9001,
  "DbHost": "localhost",
  "DbUser": "dlee",
  "DbName": "rest-api-go",
  "DbPassword": ""
}
```

A combination of the environment variable and configuration file settings will be used by the application to determine
run time settings.
