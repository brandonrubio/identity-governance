# Software Requirements Specification (SRS)

**Project Name:** RoleCheck\
**Version:** MVP\
**Date:** 2023-05-18\
**Author:** Brandon Rubio

**Table of Contents**

1. [Introduction](#introduction)
   - 1.1 [Purpose](#1-1-purpose)
   - 1.2 [Intended Audience and Use](#1-2-intended-audience-and-use)
   - 1.3 [Product Scope](#1-3-product-scope)
   - 1.4 [Definitions](#1-4-definitions)
2. [Overall Description](#overall-description)
   - 2.1 [User Roles](#2-1-user-roles)
   - 2.2 [Constraints and Assumptions](#2-2-constraints-and-assumptions)
3. [Functional Requirements](#functional-requirements)
   - 3.1 [Authentication and Authorization](#3-1-authentication-and-authorization)
   - 3.2 [Application Management](#3-2-application-management)
   - 3.3 [Role Management](#3-3-role-management)
   - 3.4 [User Assignment](#3-4-user-assignment)
   - 3.5 [Dashboard](#3-5-dashboard)
4. [Non-Functional Requirements](#non-functional-requirements)
   - 4.1 [Security](#4-1-security)
   - 4.2 [Accessibility](#4-2-accessibility)
   - 4.3 [Technology Stack](#4-3-technology-stack)
---

## Introduction <a id="introduction"></a>

### 1.1 Purpose <a id="1-1-purpose"></a>
RoleCheck is a role management platform for applications in Microsoft Entra ID. It provides a web-based dashboard for teams to manage and view user-role assignments across internal applications.

### 1.2 Intended Audience and Use <a id="1-2-intended-audience-and-use"></a>
RoleCheck is intended for teams within an organization that need to manage user-role assignments across multiple applications. This includes IT administrators, security teams, and development teams.

### 1.3 Product Scope <a id="1-3-product-scope"></a>
Features for MVP:
- Sign-in using Microsoft Entra ID
- Create new app registrations
- Create roles in each app registration
- Assign users to roles
- Web UI to manage and view relationships between applications, roles, and users

### 1.4 Definitions <a id="1-4-definitions"></a>
- **Microsoft Entra ID:** Microsoft's cloud-based identity and access management platform (formerly known as Azure Active Directory).
- **App Registration:** A registered application within Microsoft Entra ID. App registration and applications will be used interchangeably throughout this document.
- **Role:** A predefined set of permissions assigned to users within an application.
- **User:** An individual who has been assigned roles within an application.

---

## Overall Description <a id="overall-description"></a>

### 2.1 User Roles <a id="2-1-user-roles"></a>
- **Admin** Can register apps, create roles, and assign users to roles
- **Viewer** Can view applications and role assignments but cannot modify them

### 2.2 Constraints and Assumptions <a id="2-2-constraints-and-assumptions"></a>
- Self hosted and not publicly exposed
- Users are part of an Entra ID tenant
- No mobile app is needed

## Functional Requirements <a id="functional-requirements"></a>

### 3.1 Authentication and Authorization <a id="3-1-authentication-and-authorization"></a>
- **FR-1.1**: Users must sign in using Microsoft Entra ID.
- **FR-1.2**: Admin and Viewer role is assigned to a user in the Azure portal

### 3.2 Application Management <a id="3-2-application-management"></a>
- **FR-2.1**: Admins create, update, and delete application entries
- **FR-2.2**: Applications must have a unique name and optional description
- **FR-2.3**: Viewers can view application entries

### 3.3 Role Management <a id="3-3-role-management"></a>
- **FR-3.1**: Admins can create and delete roles scoped to an application
- **FR-3.2**: Roles must have a name that is unique within an application and a description

### 3.4 User Assignment <a id="3-4-user-assignment"></a>
- **FR-4.1**: Admins can search for users in Entra by email or name
- **FR-4.2**: Admins can assign or remove users from application roles
- **FR-4.3**: Only users that exist in Entra can be assigned

### 3.5 Dashboard <a id="3-5-dashboard"></a>
- **FR-5.1**: Admins and Viewers can see a list of applications
- **FR-5.2**: Each application shows a table of users assigned to roles

## Non-Functional Requirements <a id="non-functional-requirements"></a>

### 4.1 Security <a id="4-1-security"></a>
- OIDC authentication and authorization with Microsoft Entra ID.
- All endpoints require authentication
- Only Admins can mutate application data

### 4.2 Accessibility <a id="4-2-accessibility"></a>
- The UI should be usable with keyboard only (no mouse required)
- Text and buttons should have good contrast for readability
- Important actions or messages should not require on color alone
- Aria attributes should be added to all custom components

### 4.3 Technology Stack <a id="4-3-technology-stack"></a>
- **Backend**: C#, ASP.NET, PostgreSQL
- **Frontend**: TypeScript, React.js
- **Identity**: Microsoft Entra ID
- **Hosting**: Azure App Service, Azure Blob, Azure CDN, Azure Database for PostgreSQL
