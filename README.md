# Quick Start Guide

## 1. Prerequisites

- Go 1.21+ installed
- Tencent Cloud account with EdgeOne service enabled
- Domain configured in EdgeOne

## 2. Get API Credentials

1. Visit [Tencent Cloud API Keys](https://console.tencentcloud.com/cam/capi)
2. Create a new API key pair
3. Note down your `Secret ID` and `Secret Key`

## 3. Set Environment Variables

```bash
export EDGEONE_SECRET_ID="your_secret_id_here"
export EDGEONE_SECRET_KEY="your_secret_key_here"
```

## 4. Build Caddy with EdgeOne Plugin

```bash
# Using the build script
./build.sh

# Or manually with xcaddy
xcaddy build --with github.com/caddy-dns/edgeone
```

## 5. Create Configuration

Create a `Caddyfile`:

```caddyfile
{
    acme_dns edgeone {
        secret_id {env.EDGEONE_SECRET_ID}
        secret_key {env.EDGEONE_SECRET_KEY}
    }
}

your-domain.com {
    respond "Hello from Caddy with EdgeOne DNS!"
}
```

## 6. Run Caddy

```bash
./caddy run --config Caddyfile
```

That's it! Caddy will automatically obtain SSL certificates using the EdgeOne DNS challenge.

## Troubleshooting

- Ensure your domain is properly configured in EdgeOne
- Check that your API credentials have DNS management permissions
- Verify your environment variables are set correctly
- Check Caddy logs for any DNS challenge errors