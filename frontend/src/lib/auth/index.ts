/**
 * Authentication API service
 */

import { AuthResponse, SignInFormData, SignUpFormData } from '@/types/auth';

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

/**
 * Sign up user API call
 */
export const signUpUser = async (
	formData: SignUpFormData
): Promise<AuthResponse> => {
	try {
		// TODO: Change the API
		const response = await fetch(`${API_BASE_URL}/api/auth/signup`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify({
				username: formData.username,
				email: formData.email,
				phone: formData.phone,
				password: formData.password,
			}),
		});

		const data = await response.json();

		if (!response.ok) {
			throw new Error(data.message || 'Sign up failed');
		}

		return {
			success: true,
			message: 'Account created successfully',
			data: {
				userId: data.userId,
				email: data.email,
			},
		};
	} catch (error) {
		console.error('Sign up error:', error);
		return {
			success: false,
			message:
				error instanceof Error ? error.message : 'An unexpected error occurred',
		};
	}
};

/**
 * Sign in user API call
 */
export const signInUser = async (
	formData: SignInFormData
): Promise<AuthResponse> => {
	try {
		// TODO: Change the API
		const response = await fetch(`${API_BASE_URL}/api/auth/signin`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify(formData),
		});

		const data = await response.json();

		if (!response.ok) {
			throw new Error(data.message || 'Sign in failed');
		}

		return {
			success: true,
			message: 'Signed in successfully',
			data: {
				userId: data.userId,
				email: data.email,
				token: data.token,
			},
		};
	} catch (error) {
		console.error('Sign in error:', error);
		return {
			success: false,
			message: error instanceof Error ? error.message : 'Sign in failed',
		};
	}
};

/**
 * Verify email API call
 */
export const verifyEmail = async (token: string): Promise<AuthResponse> => {
	try {
		const response = await fetch(`${API_BASE_URL}/api/auth/verify-email`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
			},
			body: JSON.stringify({ token }),
		});

		const data = await response.json();

		if (!response.ok) {
			throw new Error(data.message || 'Email verification failed');
		}

		return {
			success: true,
			message: 'Email verified successfully',
		};
	} catch (error) {
		console.error('Email verification error:', error);
		return {
			success: false,
			message:
				error instanceof Error ? error.message : 'Email verification failed',
		};
	}
};

/**
 * Logout user API call
 */
export const logoutUser = async (): Promise<AuthResponse> => {
	try {
		const response = await fetch(`${API_BASE_URL}/api/auth/logout`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
			},
		});

		const data = await response.json();

		if (!response.ok) {
			throw new Error(data.message || 'Logout failed');
		}

		return {
			success: true,
			message: 'Logged out successfully',
		};
	} catch (error) {
		console.error('Logout error:', error);
		return {
			success: false,
			message: error instanceof Error ? error.message : 'Logout failed',
		};
	}
};
