import time
import requests
import logging


class LoadGenerator:
    def __init__(self, api_url, delay_ms):
        self.api_url = api_url
        # Convert milliseconds to seconds
        self.delay_ms = delay_ms / 1000.0
        self.terminate = False

    def start(self):
        """Start the load generator loop."""
        while not self.terminate:
            try:
                response = requests.get(self.api_url)
                logging.info(f"Request to {self.api_url} completed with status code {response.status_code}")
            except requests.RequestException as e:
                logging.error(f"Request to {self.api_url} failed: {e}")

            time.sleep(self.delay_ms)

        logging.info("Load generator terminated gracefully.")

    def stop(self, signum, frame):
        """Handle termination signals."""
        logging.info(f"Received signal {signum}, terminating...")
        self.terminate = True
