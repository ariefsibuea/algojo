# System Design Interview

## 1. Design a Scalable eCommerce Integration API

**Key Components:**

1. Partner onboarding
2. Order APIs
3. Asynchronous callbacks
4. Rate limiting
5. High availability

**Reference:**
[Key Considerations When Designing a Scalable API for a High-Traffic Ecommerce Platform](https://www.zigpoll.com/content/what-are-the-key-considerations-when-designing-a-scalable-api-for-a-hightraffic-ecommerce-platform)

## 2. Design a Unified Authentication & Authorization System for Partner APIs

**Challenge**: Design a system that handles authentication for thousands of partners with different access levels, supports OAuth 2.0, API keys, and webhook signatures. Must scale to millions of requests per day with sub-100ms latency.

**Key Components**:

- Identity provider service
- Token management with refresh
- Rate limiting per partner
- Permission management
- Audit logging

Notes: This directly relates to the job requirements. Focus on security, scalability, and developer experience.

**Reference:** [Authentication and Authorization](https://www.tryexponent.com/courses/system-design-interviews/authentication-authorization)
