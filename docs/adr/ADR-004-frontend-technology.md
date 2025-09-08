# ADR-004: Frontend Technology (React)

## Context
InfraVault Phase 1 requires a modern, responsive web frontend that provides:
- User authentication interface with Google OAuth 2.0
- Image upload functionality with drag-and-drop support
- Personal image gallery with CRUD operations
- Public gallery for viewing all uploaded images
- Responsive design for desktop and mobile devices
- Real-time feedback for user interactions
- Integration with the Go backend API

The frontend must be containerizable, maintainable, and provide excellent user experience while keeping development complexity appropriate for an MVP.

## Decision
We will implement the frontend using **React** with modern JavaScript tooling and a component-based architecture.

### Technology Stack:

#### Core Framework: React 18+
- **Build Tool**: Vite (fast development server and build tool)
- **Language**: TypeScript (type safety and better developer experience)
- **Styling**: Tailwind CSS (utility-first CSS framework)
- **State Management**: React Context API + useReducer (built-in state management)
- **HTTP Client**: Native `fetch`
- **Routing**: React Router
- **Form Handling**: React Hook Form (performant forms with minimal re-renders)
- **File Upload**: React Dropzone (drag-and-drop file upload)
- **UI Components**: Headless UI (unstyled, accessible UI components)
- **Icons**: Lucide Icons library

#### Development Tools:
- **Linting**: ESLint with TypeScript support
- **Formatting**: Prettier
- **Testing**: Vitest + React Testing Library
- **Type Checking**: TypeScript compiler