# Next.js Project Structure - Best Practices

This document outlines the organized file structure following Next.js best practices for the ScholarAI frontend.

## 📁 Directory Structure

```
src/
├── app/                          # Next.js App Router
│   ├── (auth)/                   # Route groups
│   │   ├── sign-in/
│   │   │   └── page.tsx
│   │   ├── sign-up/
│   │   │   └── page.tsx
│   │   └── verify-email/
│   │       └── page.tsx
│   ├── globals.css
│   ├── layout.tsx
│   └── page.tsx
├── components/                   # Reusable UI components
│   ├── auth/                     # Authentication components
│   ├── common/                   # Shared components
│   │   └── NavBar.tsx
│   └── ui/                       # UI component library
├── hooks/                        # Custom React hooks
│   ├── auth/                     # Authentication hooks
│   │   ├── useSignInForm.ts
│   │   ├── useSignUpForm.ts
│   │   └── index.ts
│   └── index.ts
├── lib/                          # Utility libraries
│   ├── auth/                     # Authentication logic
│   │   ├── index.ts              # API services
│   │   └── validation.ts         # Form validation
│   ├── validation/               # Pure validation utilities
│   │   └── index.ts
│   └── utils.ts                  # General utilities
├── stores/                       # State management
├── types/                        # TypeScript type definitions
│   ├── auth/                     # Authentication types
│   │   └── index.ts
│   └── index.ts
└── utils/                        # Utility functions
    ├── auth/
    └── validation/
```

## 🎯 Key Principles

### 1. **Separation of Concerns**
- **Components**: Pure UI rendering
- **Hooks**: State management and business logic
- **Lib**: Utility functions and API calls
- **Types**: TypeScript definitions

### 2. **Feature-Based Organization**
- Authentication features grouped together
- Related files co-located
- Clear import paths with barrel exports

### 3. **Scalability**
- Easy to add new features
- Consistent patterns across the app
- Reusable utilities and components

## 📦 Import Patterns

### Barrel Exports
```typescript
// ✅ Good - Clean imports
import { useSignUpForm, useSignInForm } from '@/hooks/auth';
import { SignUpFormData, AuthResponse } from '@/types/auth';
import { signUpUser, signInUser } from '@/lib/auth';

// ❌ Avoid - Deep imports
import { useSignUpForm } from '@/hooks/auth/useSignUpForm';
```

### Path Aliases
```typescript
// Configured in tsconfig.json
'@/components/*' -> 'src/components/*'
'@/hooks/*'      -> 'src/hooks/*'
'@/lib/*'        -> 'src/lib/*'
'@/types/*'      -> 'src/types/*'
```

## 🔧 File Responsibilities

### `/lib/auth/index.ts`
- API service functions
- Backend communication
- Error handling

### `/lib/auth/validation.ts`
- Form validation logic
- Business rules
- Error message generation

### `/lib/validation/index.ts`
- Pure validation functions
- Reusable across features
- No side effects

### `/hooks/auth/`
- Form state management
- User interaction handling
- Loading states

### `/types/auth/index.ts`
- TypeScript interfaces
- API response types
- Form data structures

## 🚀 Benefits

1. **Maintainability**: Easy to find and modify code
2. **Reusability**: Components and utilities can be shared
3. **Testability**: Clear separation makes testing easier
4. **Scalability**: New features follow established patterns
5. **Developer Experience**: Intuitive file organization

## 📝 Adding New Features

1. **Create feature folder** in appropriate directory
2. **Add barrel exports** to index files
3. **Follow naming conventions**
4. **Update type definitions**
5. **Add to documentation**

This structure follows Next.js 13+ App Router conventions and React best practices for scalable applications.
