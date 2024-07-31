#!/usr/bin

import argparse
import base64
import sys


def _parse_args() -> argparse.Namespace:
    """Build and parse command line arguments.
    Returns:
        A namespace with parsed arguemnts.
    """
    parser = argparse.ArgumentParser()
    parser.add_argument('s', help='arbitrary string to test',
                        type=str, nargs='?')
    parser.parse_args()
    return parser.parse_args()


def hex_to_base64(s) -> str:
    # The string that was originally encoded as base64.
    hex_decode = bytes.fromhex(s).decode('utf-8')
    hex_decode_bytes = bytes(hex_decode, 'utf-8')
    return base64.b64encode(hex_decode_bytes).decode('utf-8')


def run_tests():
    """Run unit tests.

    Put this in a method for deferred test suite import.
    """
    import unittest

    class TestBuildNewUser(unittest.TestCase):
        def test_convert_hex_to_base64(self) -> None:
            s = '49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d'
            r = 'SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t'
            self.assertEqual(r, hex_to_base64(s))

    suite = unittest.TestLoader().loadTestsFromTestCase(TestBuildNewUser)
    result = unittest.TextTestRunner(verbosity=2).run(suite)

    if result.errors or result.failures:
        sys.exit(1)
    else:
        sys.exit(0)


def main():
    args = _parse_args()
    if args.s is None:
        run_tests()
    else:
        hex_to_base64(args.s)

if __name__ == '__main__':
    main()
