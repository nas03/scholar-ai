/**
 * Custom hook for sign in form management
 */

import { signInUser } from '@/lib/auth';
import { hasFormErrors, validateSignInForm } from '@/lib/auth/validation';
import { FormErrors, SignInFormData } from '@/types/auth';
import { useState } from 'react';

export const useSignInForm = () => {
	const [formData, setFormData] = useState<SignInFormData>({
		email: '',
		password: '',
	});

	const [errors, setErrors] = useState<Partial<FormErrors>>({});
	const [isSubmitting, setIsSubmitting] = useState(false);

	/**
	 * Handle input change and clear errors
	 */
	const handleInputChange =
		(field: keyof SignInFormData) =>
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
		const newErrors = validateSignInForm(formData);
		setErrors(newErrors);
		return !hasFormErrors(newErrors as FormErrors);
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
			const result = await signInUser(formData);

			if (result.success) {
				// Handle successful sign in
				console.log('Sign in successful:', result);
				// You might want to redirect to dashboard or store token
				// router.push('/dashboard');
			} else {
				// Handle sign in error
				console.error('Sign in failed:', result.message);
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
			email: '',
			password: '',
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
