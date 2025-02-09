import time

def sqrt(x: float) -> float:
    start_time = time.time()  # Start time measurement
    z = 1.0  # Initial guess
    
    print(f"Sqrt of {x}")
    for i in range(1, 11):  # 10 iterations
        z -= (z * z - x) / (2 * z)  # Newton's method
        print(f"Iteration {i}: {z}")

    elapsed_time = time.time() - start_time  # End time measurement
    print(f"Execution time: {elapsed_time:.10f} seconds")

    return z

def main():
    print(sqrt(0.125), sqrt(2), sqrt(4), sqrt(16))

if __name__ == "__main__":
    main()
