'use client';
import { useSignUpForm } from '@/hooks/auth';
import {
	Box,
	Button,
	CircularProgress,
	FormControl,
	FormHelperText,
	Input,
	InputLabel,
	Typography,
} from '@mui/material';
import Link from 'next/link';

const SignUpPage = () => {
	const {
		formData,
		errors,
		isSubmitting,
		handleInputChange,
		handleSubmit,
	} = useSignUpForm();

	return (
		<>
			<Box
				component="form"
				onSubmit={handleSubmit}
				className="flex flex-col gap-3 w-1/3 relative left-1/2 -translate-x-1/2 top-[20rem] justify-center border-2 p-5">
				<Typography
					variant="h3"
					gutterBottom
					className="text-center text-blue-500 font-semibold">
					ScholarAI
				</Typography>

				<FormControl variant="standard" error={!!errors.username}>
					<InputLabel htmlFor="username">Username</InputLabel>
					<Input
						id="username"
						value={formData.username}
						onChange={handleInputChange('username')}
						required
					/>
					{errors.username && (
						<FormHelperText>{errors.username}</FormHelperText>
					)}
				</FormControl>

				<FormControl variant="standard" error={!!errors.email}>
					<InputLabel htmlFor="email">Email</InputLabel>
					<Input
						type="email"
						id="email"
						value={formData.email}
						onChange={handleInputChange('email')}
						required
					/>
					{errors.email && <FormHelperText>{errors.email}</FormHelperText>}
				</FormControl>

				<FormControl variant="standard" error={!!errors.phone}>
					<InputLabel htmlFor="phone">Phone Number</InputLabel>
					<Input
						type="tel"
						id="phone"
						value={formData.phone}
						onChange={handleInputChange('phone')}
						required
					/>
					{errors.phone && <FormHelperText>{errors.phone}</FormHelperText>}
				</FormControl>

				<FormControl variant="standard" error={!!errors.password}>
					<InputLabel htmlFor="password">Password</InputLabel>
					<Input
						type="password"
						id="password"
						value={formData.password}
						onChange={handleInputChange('password')}
						required
					/>
					{errors.password && (
						<FormHelperText>{errors.password}</FormHelperText>
					)}
				</FormControl>

				<FormControl variant="standard" error={!!errors.confirmPassword}>
					<InputLabel htmlFor="confirm-password">Confirm Password</InputLabel>
					<Input
						type="password"
						id="confirm-password"
						value={formData.confirmPassword}
						onChange={handleInputChange('confirmPassword')}
						required
					/>
					{errors.confirmPassword && (
						<FormHelperText>{errors.confirmPassword}</FormHelperText>
					)}
				</FormControl>

				<Button 
					type="submit" 
					variant="contained" 
					fullWidth
					disabled={isSubmitting}
					startIcon={isSubmitting ? <CircularProgress size={20} /> : null}
				>
					{isSubmitting ? 'Creating Account...' : 'Sign up'}
				</Button>

				<p className="text-center">
					Already got an account?{' '}
					<Link
						href={'/sign-in'}
						className="text-blue-500 font-semibold hover:underline">
						Sign in
					</Link>
				</p>
			</Box>
		</>
	);
};

export default SignUpPage;
