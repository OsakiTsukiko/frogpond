import os
import base64

def generate_secret_key():
    # Generate a 256-bit (32-byte) random key
    key = os.urandom(32)
    # Encode the key in Base64 for easy storage and sharing
    return base64.urlsafe_b64encode(key).decode('utf-8')

if __name__ == "__main__":
    secret_key = generate_secret_key()
    print(f"Generated JWT Secret Key: {secret_key}")