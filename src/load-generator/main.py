import os
import signal
import sys
import logging
from HTTPGenerator import LoadGenerator

if __name__ == "__main__":
    # Configure logging
    logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')

    # Read environment variables
    # Default to 1000 milliseconds if not set
    api_url = os.getenv('API_URL')
    delay_ms = float(os.getenv('DELAY_MS', 1000))

    if not api_url:
        logging.error("Error: API_URL environment variable is not set.")
        sys.exit(1)

    # Instantiate LoadGenerator
    generator = LoadGenerator(api_url, delay_ms)

    # Set up signal handlers
    signal.signal(signal.SIGTERM, generator.stop)
    signal.signal(signal.SIGINT, generator.stop)

    logging.info(f"Starting load generator for {api_url} with a delay of {delay_ms} milliseconds between requests.")
    generator.start()
