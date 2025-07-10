
# 📗 IT SCADA System Functional Module Design Document (Final Version)

## 1️⃣ System Overview

**System Purpose:**

To build a comprehensive IT equipment management platform with the following features:

- Real-time monitoring of equipment status
- Automated alerting and event handling
- Complete asset and work order management
- Centralized management of master data (factories, departments, personnel)

**Technology Stack:**

- Frontend: React.js
- Backend: Node.js / Typescript / express / swagger
- Database: MongoDB
- API: RESTful API
- Real-time Communication: WebSocket

## 2️⃣ Module Architecture Diagram

```mermaid
graph LR
A[Master Data Module (Factories, Departments, Personnel)] --> D[Asset Management Module]
D --> C[Event and Work Order Management Module]
B[Alert Module] --> C
E[Real-Time Monitoring Module] --WebSocket Notification--> B
```

## 3️⃣ Module Details

### 📋 Master Data Module
- Provides basic company data settings (including factories, departments/units, personnel information)
- Source of base data for all assets, work orders, and dispatching
- Provides personnel role management and basic authorization information

### 🖥️ Real-Time Monitoring Module
- Real-time display of device health status
- Uses WebSocket to push status updates instantly

### 🚨 Alert Module
- Automatically detects anomalies and notifies the frontend in real time
- Automatically triggers and creates event work orders

### 🛠️ Event and Work Order Management Module
- Automatically records and manages equipment anomaly events
- Supports creation and assignment of work orders, recording maintenance content
- Work order process statuses: Pending → Assigned → In Progress → Completed → Closed

### 📦 Asset Management Module (CMDB)
- Records basic IT equipment information (including location, warranty, responsible department)
- Manages equipment maintenance and change history
- Provides device-to-personnel and unit binding information

## 4️⃣ Example API Design

| Method | Path                   | Description                |
|--------|------------------------|----------------------------|
| GET    | `/api/factories`       | Get factory list           |
| POST   | `/api/factories`       | Create new factory         |
| GET    | `/api/units`           | Get department list        |
| POST   | `/api/units`           | Create new department      |
| GET    | `/api/users`           | Get personnel list         |
| POST   | `/api/users`           | Create new personnel       |
| GET    | `/api/assets`          | Get asset list             |
| POST   | `/api/assets`          | Create new asset           |
| GET    | `/api/events`          | Get event and work order list |
| POST   | `/api/events`          | Create new event/work order   |
| PATCH  | `/api/events/:id`      | Update event/work order    |

## 5️⃣ Data Structure Examples (MongoDB Schema)

```json
// Factory Schema
{ "_id": ObjectId, "name": String, "address": String, "departments": [ObjectId] }

// Unit Schema
{ "_id": ObjectId, "name": String, "factory_id": ObjectId, "manager": ObjectId }

// User Schema
{ "_id": ObjectId, "name": String, "email": String, "phone": String, "role": String, "unit_id": ObjectId }
```
## 🔢 Recommended Module Development Sequence

1. **Master Data Module** (Factory, Department, Personnel)
2. **Asset Management Module** (CMDB)
3. **Event and Work Order Management Module**
4. **Real-Time Monitoring Module**
5. **Alert Module**

## 6️⃣ Appendix

- Provide system architecture diagrams, data flow diagrams, and UI mockups for team reference.
