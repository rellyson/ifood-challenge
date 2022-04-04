from signal import SIGINT, SIGTERM, signal


class SignalHandler:
    def __init__(self):
        self.received_signal = False
        signal(SIGINT, self._signal_handler)
        signal(SIGTERM, self._signal_handler)

    def _signal_handler(self, signal):
        print("Received signal {}, exiting gracefully...".format(signal))
        self.received_signal = True
