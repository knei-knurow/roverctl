import tkinter as tk

window = tk.Tk()

window.rowconfigure(index=0, weight=1, minsize=50)
window.columnconfigure(index=0, weight=1, minsize=75)

frame = tk.Frame(master=window, relief=tk.RAISED, borderwidth=1)
frame.grid(row=0, column=0, padx=5, pady=5)

label = tk.Label(master=frame, text=f"GO commands")
label.pack(padx=5, pady=5)
label = tk.Label(master=frame, text=f"WARNING: \nalways stop when changing direction!")
label.pack(padx=5, pady=5)

window.rowconfigure(index=1, weight=1, minsize=50)
window.columnconfigure(index=0, weight=1, minsize=75)

frame = tk.Frame(master=window, relief=tk.RAISED, borderwidth=1)
frame.grid(row=1, column=0, padx=5, pady=5)

btnGoForward = tk.Button(master=frame, text="go forward")
btnGoForward.pack(padx=5, pady=5)


btnGoStop = tk.Button(master=frame, text="go stop")
btnGoStop.pack(padx=5, pady=5)


btnGoBackward = tk.Button(master=frame, text="go backward")
btnGoBackward.pack(padx=5, pady=5)

# for i in range(3):
#     window.rowconfigure(index=i, weight=1, minsize=50)
#     window.columnconfigure(index=i, weight=1, minsize=75)

#     for j in range(0, 3):
#         frame = tk.Frame(master=window, relief=tk.RAISED, borderwidth=1)
#         frame.grid(row=i, column=j, padx=5, pady=5)

#         label = tk.Label(master=frame, text=f"Row {i}\nColumn {j}")
#         label.pack(padx=5, pady=5)

window.mainloop()
