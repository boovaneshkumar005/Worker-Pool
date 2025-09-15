Worker Pool in Go – A Simple Banking Example 🏦

This little project is a hands-on demo of the worker pool concept in Go.
Instead of using boring examples, I’ve framed it around a bank deposit system so it feels more real.

What’s Happening Here?

Imagine you run a bank.

Customers walk in and make deposits.

You have a few cashiers (workers) who handle these deposits.

Every time a cashier processes a deposit, they deduct a flat fee.

At the end of the day, you put all transactions into a passbook and calculate the total balance.

That’s exactly what this program does, using goroutines, channels, and WaitGroups.

How It Works (Step by Step)

We create 10 random deposits (like 10 customers walking in).

We start 5 workers (like 5 bank clerks).

Deposits go into a channel (think of it as a queue).

Workers pick deposits from the queue, deduct a fee, and return the result.

We collect all results into a passbook and print everything out neatly.

Running the Code
go run main.go


You’ll see something like this:

📒 Passbook Records:
Customer: Cust-1 | Amount: 534
Customer: Cust-2 | Amount: 210
Customer: Cust-3 | Amount: 743
Customer: Cust-4 | Amount: 125
Customer: Cust-5 | Amount: 863
...
Total Balance: 4210

Why This Matters

Worker pools help you handle a lot of tasks at once without wasting resources.

Instead of spinning up a goroutine for every single task (which can crash your app if tasks are huge), you create a fixed set of workers that handle jobs efficiently.

This is useful in real projects like:

Processing API requests

Handling background jobs

Image/file uploads

Any situation with lots of tasks that can be done in parallel
