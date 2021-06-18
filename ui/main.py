import tkinter as tk
import subprocess

window = tk.Tk()

speedVar = tk.IntVar()
speedVar.initialize(120)

degreesVar = tk.IntVar()
degreesVar.initialize(90)

# handlers
def handle_go_forward(event: tk.Event):
    print(f"handle_go")
    proc = subprocess.run(
        [
            "./roverctl",
            "go",
            "--speed",
            f"{speedVar.get()}",
            "forward",
        ]
    )
    print(f"return code: {proc.returncode}")


def handle_go_stop(event: tk.Event):
    print(f"handle_go_stop")
    proc = subprocess.run(
        [
            "./roverctl",
            "go",
            "--speed",
            "255",
            "stop",
        ]
    )
    print(f"return code: {proc.returncode}")


def handle_go_backward(event: tk.Event):
    print(f"handle_go_backward")
    proc = subprocess.run(
        [
            "./roverctl",
            "go",
            "--speed",
            f"{speedVar.get()}",
            "backward",
        ]
    )
    print(f"return code: {proc.returncode}")


def handle_turn(event: tk.Event):
    print(f"handle_go_turn")
    proc = subprocess.run(
        [
            "./roverctl",
            "turn",
            "--degrees",
            f"{degreesVar.get()}",
        ]
    )
    print(f"return code: {proc.returncode}")


# window.rowconfigure(index=0, weight=1, minsize=50)
# window.columnconfigure(index=0, weight=1, minsize=75)

frame0 = tk.Frame(master=window, relief=tk.RAISED, borderwidth=1)
frame0.grid(row=0, column=0, padx=5, pady=5)

label = tk.Label(master=frame0, text=f"GO commands")
label.pack(padx=5, pady=5)
label = tk.Label(master=frame0, text=f"WARNING: \nalways stop when changing direction!")
label.pack(padx=5, pady=5)

labelSpeed = tk.Label(master=frame0, text=f"speed")
labelSpeed.pack(padx=5, pady=5)

entryDegrees = tk.Entry(
    master=frame0,
    textvariable=speedVar,
    fg="black",
    width=10,
    font=("calibre", 10, "normal"),
)
entryDegrees.pack()

btnGoForward = tk.Button(master=frame0, text="go forward")
btnGoForward.pack(padx=5, pady=5)
btnGoForward.bind("<Button-1>", handle_go_forward)


btnTurn = tk.Button(master=frame0, text="go stop")
btnTurn.pack(padx=5, pady=5)
btnTurn.bind("<Button-1>", handle_go_stop)


btnGoBackward = tk.Button(master=frame0, text="go backward")
btnGoBackward.pack(padx=5, pady=5)
btnGoBackward.bind("<Button-1>", handle_go_backward)

label = tk.Label(master=frame0, text=f"TURN commands")
label.pack(padx=5, pady=5)

degreesSpeed = tk.Label(master=frame0, text=f"degrees")
degreesSpeed.pack(padx=5, pady=5)

entryDegrees = tk.Entry(
    master=frame0,
    textvariable=degreesVar,
    fg="black",
    width=10,
    font=("calibre", 10, "normal"),
)
entryDegrees.pack()

btnTurn = tk.Button(master=frame0, text="turn")
btnTurn.pack(padx=5, pady=5)
btnTurn.bind("<Button-1>", handle_turn)


window.mainloop()
