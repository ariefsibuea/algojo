# Parking Lot

**Context:** Simulated 35-minute low-level design session for a Backend Engineer role at a pay-later fintech. Unassigned-spot variant — modeled on Secure Parking / ISS-style operations common in Indonesian malls and office buildings. I play both Interviewer and Candidate. Code in Go.

## Requirements

### Q&A Simulation

- **Clarifying question**

  > **Candidate:** "To confirm — does the system assign a specific spot at entry, or does it only issue a timed ticket and let the driver pick any free spot in the correct zone?"
  >
  > **Interviewer:** "Second one. Ticket at the gate, driver finds their own spot. The system only tracks aggregate occupancy per vehicle type."

- **On vehicle types and zones**

  > **Candidate:** "What vehicle types and what zoning? I'm assuming motorcycle and car, with separate physical areas and separate capacity counters."
  >
  > **Interviewer:** "Correct. Motorcycle and car only. Physically separate zones — we don't need to enforce that in software, the gates and signage handle it. But we do track capacity per type."

- **On gates and concurrency**

  > **Candidate:** "Single entry gate and single exit gate, or multiple? If multiple, do I need to worry about two cars arriving at different entry gates at the same time and both being told the lot has one slot left?"
  >
  > **Interviewer:** "Multiple gates — assume 2 entry, 2 exit. You should handle concurrent access correctly. No overselling."

- **On fee structure**

  > **Candidate:** "Fee calculation — tiered hourly? Daily cap? Any grace period?"
  >
  > **Interviewer:** "Tiered hourly, different rates per vehicle type, first 15 minutes free (grace period), and there's a 24-hour flat cap so one ticket never exceeds one day's rate no matter how long it sits."

- **On verification**

  > **Candidate:** "How do we prevent ticket fraud — say, someone handing a motorcycle ticket to a car driver? Do we record and verify the plate?"
  >
  > **Interviewer:** "Yes. Plate is captured at entry. On exit, the plate presented must match the ticket. Mismatch → reject."

- **On lost tickets**

  > **Candidate:** "What happens if someone loses their ticket? In most Indonesian lots I've seen, there's a flat penalty."
  >
  > **Interviewer:** "Good, you've parked here. Flat penalty equal to the 24-hour cap, plus the plate must still match something we issued today."

- **On scope**
  > **Candidate:** "Out of scope: actual payment rails, gate hardware / IoT, UI, persistence, camera-based plate recognition — just the logical plate string as input?"
  >
  > **Interviewer:** "Yes. All out of scope."

### Final Requirements

1. Two vehicle types: motorcycle and car. Each has its own zone with a fixed capacity counter.
2. No spot assignment. Driver picks freely within the correct zone.
3. Multiple gates: 2 entry, 2 exit.
4. Enter gate(vehicle) -> Ticket, error
   - Ticket carries: ID, plate, vehicle type, entry time
   - Reject if zone counter has no remaining capacity
5. Exit gate(ticketID, plate) -> fee, error
   - Verify ticket exist
   - Verify plate matches the ticket
   - Compute free from entry/exit time + vehicle type
   - Release counter
6. Lost ticket(plate, vehicleType) -> fee, error
   - Charge flat penalty = 24-hour cap
   - Plate must match at least one active ticket issued today
   - Release the counter
7. Get available slot(vehicle type) -> int
8. Fee calculation:
   - First 15 minutes free
   - Tiered hourly rate per vehicle type (first hours + subsequent hour)
   - Capped at a daily flat rate regardless of duration

Out of Scope:

1. Payment processing
2. Persistence
3. Gate hardware, IoT, camera
4. UI / API layer
5. Reservation
6. Multi-building

## Entities and Relationships

The model is genuinely lean here. Scanning nouns:

| Entity                           | Why it exists                                                                                                                       |
| -------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- |
| **ParkingLot**                   | Orchestrator. Holds zones, active tickets, rate calculator. Entry point for all gate operations.                                    |
| **Zone**                         | One per vehicle type. Holds capacity and current occupied count. **Not** a list of spots — just an integer counter with invariants. |
| **Ticket**                       | Issued on entry, consumed on exit. Carries ID, plate, vehicle type, entry time.                                                     |
| **Vehicle**                      | Plate + type. Value object used at entry.                                                                                           |
| **RateCalculator** _(interface)_ | Pricing strategy — tiered hourly, daily cap, grace period. Same reasoning as before: pricing rules churn, keep them swappable.      |

Notice what's **not here**: no `ParkingSpot`, no per-level free pool, no allocation algorithm. The whole allocation subsystem evaporated because the human driver is doing that job now.

Relationships:

```
ParkingLot ──owns──> map[VehicleType]*Zone
ParkingLot ──holds──> map[ticketID]*Ticket     (active tickets)
ParkingLot ──uses──>  RateCalculator
Ticket     ──refs──>  VehicleType, plate, entry time
```
