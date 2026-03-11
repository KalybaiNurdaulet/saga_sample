SAGA PATTERN IMPLEMENTATION WITHIN A SINGLE MICROSERVICE (CHECKOUT WORKFLOW)
PROJECT OVERVIEW

This project demonstrates a programmatic implementation of the Saga pattern
using the Go programming language to manage a transactional checkout workflow.
Within a single microservice, the pattern is utilized to ensure data
consistency across multiple domain areas (Payment, Inventory, and Shipping)
without relying on heavy distributed database transactions.

The implementation follows an Orchestration-based Saga approach, where a
central component manages the sequence of steps and initiates compensating
actions in the event of failure.

ARCHITECTURAL DECISIONS

The project is structured into logical layers following Clean Architecture
principles:

"saga" Package (Core Engine): Contains abstractions and a universal execution
mechanism. It is completely decoupled from business logic and can be reused
for other transactional processes.

"service" Package (Domain Logic): Contains implementations of specific
business steps (PaymentStep, InventoryStep, ShippingStep). Each implementation
encapsulates execution logic (Execute) and reversal logic (Compensate).

Entry Point (cmd/checkout): Initializes dependencies and runs the workflow
emulation.

COMPENSATION MECHANISM

The orchestrator's logic ensures the following conditions are met:

Steps are executed in a strict sequential order.

Upon successful execution, a step is added to the stack of completed actions.

If an error occurs at any stage, execution is immediately halted.

The orchestrator iterates through the stack of completed actions in reverse
order (LIFO - Last-In-First-Out) and invokes the Compensate method for
each finished step.

Compensation is not invoked for the step that caused the failure, as it is
assumed the action did not complete successfully.

PROJECT STRUCTURE

internal/saga/
step.go - Step interface with Name, Execute, and Compensate methods.
orchestrator.go - Logic for transaction management and rollback handling.
internal/service/
checkout_saga.go - High-level service for assembling and running the saga.
payment.go - Domain logic for payment processing.
inventory.go - Domain logic for inventory reservation.
shipping.go - Domain logic for logistics and shipping.
cmd/checkout/
main.go - Application entry point and scenario demonstration.

RUNNING THE APPLICATION

To run the emulation, ensure that Go version 1.20+ is installed.

Command to run:
go run cmd/checkout/main.go

During execution, logs will be sent to standard output (stdout), demonstrating
a successful end-to-end flow as well as a failure scenario with an automatic
rollback of previously completed actions.

TECHNICAL HIGHLIGHTS

Interface-Driven Design: Ensures low coupling between components and
simplifies Unit Testing through the use of mock objects.

Context Management: All methods accept context.Context, allowing for proper
handling of timeouts and cancellation signals throughout the lifecycle.

Reverse Order Execution: Implementing rollbacks in reverse order minimizes
risks associated with resource dependencies and race conditions.
