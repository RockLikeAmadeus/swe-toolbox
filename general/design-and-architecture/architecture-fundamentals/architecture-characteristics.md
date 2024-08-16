# Architecture Characteristics

**Important Note: Applications can support a large number of these architecture characterists, but _shouldn't_, as supporting a specific characteristic adds complexity to the design. The first part of defining a software architecture is deciding _which_ characteristics to support, and the aim should be to support the fewest possible, rather than the most.**

# Partial List

## Operational Characteristics

### Availability

How long is the system available (likely 24/7)? Highly available systems need to be able to go from cold start to functional quickly, in case of any failures.

### Continuity

Disaster-Recovery Capability

###  Performance

Includes stress-testing, peak analysis, analysis of the frequency of functions used, capacity required, and response times.

### Recoverability

Continuity requirements; see Continuity and Availability. This will affect backup strategy and requirements for duplicated hardware.

### Reliability and Safety

Is it mission critical, especially in a way that affects lives?

### Robustness

Ability to handle error and boundary conditions while running, or in the event of hardware or network failures.

### Scalability

Ability for the system to continue to operate as the number of users or requests increases.

## Structural Characteristics

### Configurability

Ability for end users to easily change aspects of the system's behavior through usable interfaces.

### Extensibility

How important it is to be able to "plug in" new features and functionality.

### Installability

Ease of system installation.

### Leverageability and Reuse

Ability to leverage common components across multiple products.

### Localization

Support for multiple languages.

### Maintainability

How easy is it to support and apply changes to the system?

### Portability

Does the system need to run on more than one platform?

### Supportability

What level of technical support is needed by the application? What level of logging and other facilities are required to debug errors in the system?

### Upgradeability

Ability to easily and quickly upgrade from previous versions to new versions (on servers and clients)

## Cross-Cutting Characteristics

### Accessibility

Access to all potential users, including those with disabilities.

### Archivability

Will the data need to be archived? What is the lifetime of the data?

### Authentication

Security requirements to ensure users are who they say they are.

### Authorization

Security requirements to ensure users can only access the application functionality for which they are priveleged, by use case, subsystem, web page, business rule, field level, and so on.

### Legal

Under what legal constraints is the system operating?

### Privacy

Ability to hide user data even from internal system maintainers.

### Security

Does the system data need to be encrypted either at rest, in-transit, or both? What type of authentication needs to be in place for remote access?

### Usability and Achieveability

What level of training, if any, is needed for the system to be functional with new users?