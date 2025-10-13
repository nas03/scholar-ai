/**
 * Authentication validation utilities
 */

export interface ValidationResult {
  isValid: boolean;
  message?: string;
}

/**
 * Validates email format
 */
export const validateEmail = (email: string): ValidationResult => {
  if (!email.trim()) {
    return { isValid: false, message: 'Email is required' };
  }
  
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  if (!emailRegex.test(email)) {
    return { isValid: false, message: 'Please enter a valid email address' };
  }
  
  return { isValid: true };
};

/**
 * Validates phone number format
 */
export const validatePhone = (phone: string): ValidationResult => {
  if (!phone.trim()) {
    return { isValid: false, message: 'Phone number is required' };
  }
  
  // Updated regex to handle international formats better
  const phoneRegex = /^[\+]?[1-9][\d]{0,15}$/;
  const cleanPhone = phone.replace(/\s/g, '');
  
  if (!phoneRegex.test(cleanPhone)) {
    return { isValid: false, message: 'Please enter a valid phone number' };
  }
  
  return { isValid: true };
};

/**
 * Validates password strength
 */
export const validatePassword = (password: string): ValidationResult => {
  if (!password) {
    return { isValid: false, message: 'Password is required' };
  }
  
  if (password.length < 6) {
    return { isValid: false, message: 'Password must be at least 6 characters' };
  }
  
  if (password.length > 128) {
    return { isValid: false, message: 'Password must be less than 128 characters' };
  }
  
  return { isValid: true };
};

/**
 * Validates username
 */
export const validateUsername = (username: string): ValidationResult => {
  if (!username.trim()) {
    return { isValid: false, message: 'Username is required' };
  }
  
  if (username.length < 3) {
    return { isValid: false, message: 'Username must be at least 3 characters' };
  }
  
  if (username.length > 30) {
    return { isValid: false, message: 'Username must be less than 30 characters' };
  }
  
  // Check for valid characters (alphanumeric, underscore, hyphen)
  const usernameRegex = /^[a-zA-Z0-9_-]+$/;
  if (!usernameRegex.test(username)) {
    return { isValid: false, message: 'Username can only contain letters, numbers, underscores, and hyphens' };
  }
  
  return { isValid: true };
};

/**
 * Validates password confirmation
 */
export const validatePasswordConfirmation = (password: string, confirmPassword: string): ValidationResult => {
  if (!confirmPassword) {
    return { isValid: false, message: 'Please confirm your password' };
  }
  
  if (password !== confirmPassword) {
    return { isValid: false, message: 'Passwords do not match' };
  }
  
  return { isValid: true };
};
