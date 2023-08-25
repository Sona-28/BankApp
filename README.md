# BankAppDemo

## Resource locking
* In go lang, we will use the concept of mutex to lock the resources
    1. Mutex.lock
    2. Mutex.unlock

## Steps to create transactons in mongo 
1. Create the session from mongo client
2. Start a transction with session that we created in step 1

## Extend the code with transaction
1. Creating a model called transaction - id, from, to, amount, stamp 
    Account - id, balance
2. Balance field in customer
3. entity - transaction - id, from, to, amount, timestamp 
4. Service method - tranfer money in transaction -parameters(from, to, amount)