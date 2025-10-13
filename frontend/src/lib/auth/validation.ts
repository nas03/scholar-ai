/**
 * Form validation handlers
 */

import { validateEmail, validatePassword, validatePasswordConfirmation, validatePhone, validateUsername } from '@/lib/validation';
import { FormErrors, SignInFormData, SignUpFormData } from '@/types/auth';

/**
 * Form validation handler for sign up
 */
export const validateSignUpForm = (formData: SignUpFormData): FormErrors => {
  const errors: FormErrors = {};
  
  // Validate username
  const usernameValidation = validateUsername(formData.username);
  if (!usernameValidation.isValid) {
    errors.username = usernameValidation.message;
  }
  
  // Validate email
  const emailValidation = validateEmail(formData.email);
  if (!emailValidation.isValid) {
    errors.email = emailValidation.message;
  }
  
  // Validate phone
  const phoneValidation = validatePhone(formData.phone);
  if (!phoneValidation.isValid) {
    errors.phone = phoneValidation.message;
  }
  
  // Validate password
  const passwordValidation = validatePassword(formData.password);
  if (!passwordValidation.isValid) {
    errors.password = passwordValidation.message;
  }
  
  // Validate password confirmation
  const confirmPasswordValidation = validatePasswordConfirmation(formData.password, formData.confirmPassword);
  if (!confirmPasswordValidation.isValid) {
    errors.confirmPassword = confirmPasswordValidation.message;
  }
  
  return errors;
};

/**
 * Form validation handler for sign in
 */
export const validateSignInForm = (formData: SignInFormData): Partial<FormErrors> => {
  const errors: Partial<FormErrors> = {};
  
  // Validate email
  const emailValidation = validateEmail(formData.email);
  if (!emailValidation.isValid) {
    errors.email = emailValidation.message;
  }
  
  // Validate password
  const passwordValidation = validatePassword(formData.password);
  if (!passwordValidation.isValid) {
    errors.password = passwordValidation.message;
  }
  
  return errors;
};

/**
 * Check if form has any errors
 */
export const hasFormErrors = (errors: FormErrors): boolean => {
  return Object.values(errors).some(error => error !== undefined);
};
