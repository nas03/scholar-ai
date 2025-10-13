'use client';
import { useSignInForm } from '@/hooks/auth';
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

const SignInPage = () => {
	const {
		formData,
		errors,
		isSubmitting,
		handleInputChange,
		handleSubmit,
	} = useSignInForm();

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

				<FormControl variant="standard" error={!!errors.email}>
					<InputLabel htmlFor="email">Email</InputLabel>
					<Input
						type="email"
						id="email"
						value={formData.email}
						onChange={handleInputChange('email')}
						required
					/>
					{errors.email && (
						<FormHelperText>{errors.email}</FormHelperText>
					)}
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

				<Button 
					type="submit" 
					variant="contained" 
					fullWidth
					disabled={isSubmitting}
					startIcon={isSubmitting ? <CircularProgress size={20} /> : null}
				>
					{isSubmitting ? 'Signing In...' : 'Sign In'}
				</Button>

				<p className="text-center">
					Do not have an account?{' '}
					<Link href={'/sign-up'} className="text-blue-500 font-semibold hover:underline">
						Create account
					</Link>
				</p>
			</Box>
		</>
	);
};

export default SignInPage;
