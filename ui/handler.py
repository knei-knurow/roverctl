import tkinter as tk

import subprocess


def handle_go(event: tk.Event):
    print(f"handle_go")
    proc = subprocess.run(["./roverctl", "go", "--speed", "120", "forward"])
    print(f"return code: {proc.returncode}")


def handle_go_stop(event: tk.Event):
    print(f"handle_go_stop")
    subprocess.run(["./roverctl", "go", "--speed", "255", "stop"])


def handle_turn(degrees: int):
    print(f"handle_turn: {degrees=}")
