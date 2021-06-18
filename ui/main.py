import tkinter as tk
import subprocess

window = tk.Tk()

speedVar = tk.IntVar()
degreesVar = tk.IntVar()

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
    print(f"handle_go")
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

# window.rowconfigure(index=1, weight=1, minsize=50)
# window.columnconfigure(index=0, weight=1, minsize=75)

# entrySpeed = tk.Entry(master=frame1, textvariable=speedVar, fg="red", width=10)

btnGoForward = tk.Button(master=frame0, text="go forward")
btnGoForward.pack(padx=5, pady=5)
btnGoForward.bind("<Button-1>", handle_go_forward)


btnTurn = tk.Button(master=frame0, text="go stop")
btnTurn.pack(padx=5, pady=5)
btnTurn.bind("<Button-1>", handle_go_stop)


btnGoBackward = tk.Button(master=frame0, text="go backward")
btnGoBackward.pack(padx=5, pady=5)
btnGoBackward.bind("<Button-1>", handle_go_backward)

btnTurn = tk.Button(master=frame0, text="turn")
btnTurn.pack(padx=5, pady=5)
btnTurn.bind("<Button-1>", handle_turn)

# frame0 = tk.Frame(master=window, relief=tk.RAISED, borderwidth=1)
# frame0.grid(row=0, column=1, padx=5, pady=5)


window.mainloop()
