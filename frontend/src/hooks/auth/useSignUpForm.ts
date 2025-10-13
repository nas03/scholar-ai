/**
 * Custom hook for sign up form management
 */

import { signUpUser } from '@/lib/auth';
import { hasFormErrors, validateSignUpForm } from '@/lib/auth/validation';
import { FormErrors, SignUpFormData } from '@/types/auth';
import { useState } from 'react';

export const useSignUpForm = () => {
	const [formData, setFormData] = useState<SignUpFormData>({
		username: '',
		email: '',
		phone: '',
		password: '',
		confirmPassword: '',
	});

	const [errors, setErrors] = useState<FormErrors>({});
	const [isSubmitting, setIsSubmitting] = useState(false);

	/**
	 * Handle input change and clear errors
	 */
	const handleInputChange =
		(field: keyof SignUpFormData) =>
		(event: React.ChangeEvent<HTMLInputElement>) => {
			const value = event.target.value;
			setFormData((prev) => ({ ...prev, [field]: value }));

			// Clear error when user starts typing
			if (errors[field]) {
				setErrors((prev) => ({ ...prev, [field]: undefined }));
			}
		};

	/**
	 * Validate form and set errors
	 */
	const validateForm = (): boolean => {
		const newErrors = validateSignUpForm(formData);
		setErrors(newErrors);
		return !hasFormErrors(newErrors);
	};

	/**
	 * Handle form submission
	 */
	const handleSubmit = async (event: React.FormEvent) => {
		event.preventDefault();

		if (!validateForm()) {
			return;
		}

		setIsSubmitting(true);

		try {
			const result = await signUpUser(formData);

			if (result.success) {
				// Handle successful sign up
				console.log('Sign up successful:', result);
				// You might want to redirect to verification page or show success message
				// router.push('/verify-email');
			} else {
				// Handle sign up error
				console.error('Sign up failed:', result.message);
				// You might want to show error message to user
			}
		} catch (error) {
			console.error('Unexpected error:', error);
		} finally {
			setIsSubmitting(false);
		}
	};

	/**
	 * Reset form to initial state
	 */
	const resetForm = () => {
		setFormData({
			username: '',
			email: '',
			phone: '',
			password: '',
			confirmPassword: '',
		});
		setErrors({});
		setIsSubmitting(false);
	};

	return {
		formData,
		errors,
		isSubmitting,
		handleInputChange,
		handleSubmit,
		resetForm,
	};
};