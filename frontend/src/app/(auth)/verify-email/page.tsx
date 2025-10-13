'use client';
import { OTPInput } from '@/components/auth';
import { useOTPVerification } from '@/hooks/auth';
import {
	Alert,
	Box,
	Button,
	CircularProgress,
	Typography,
} from '@mui/material';
import Link from 'next/link';

const VerifyEmailPage = () => {
	const {
		otp,
		isVerifying,
		error,
		isVerified,
		handleOTPChange,
		handleOTPComplete,
		resendOTP,
	} = useOTPVerification();

	if (isVerified) {
		return (
			<Box className="flex flex-col gap-4 w-1/3 relative left-1/2 -translate-x-1/2 top-[20rem] justify-center border-2 p-5">
				<Alert severity="success" sx={{ mb: 2 }}>
					Email verified successfully!
				</Alert>
				<Typography variant="h4" className="text-center text-green-600 font-semibold">
					Welcome to ScholarAI
				</Typography>
				<Typography variant="body1" className="text-center text-gray-600">
					Your email has been verified. You can now access all features.
				</Typography>
				<Button
					variant="contained"
					fullWidth
					component={Link}
					href="/sign-in"
					sx={{ mt: 2 }}
				>
					Continue to Sign In
				</Button>
			</Box>
		);
	}

	return (
		<Box className="flex flex-col gap-4 w-1/3 relative left-1/2 -translate-x-1/2 top-[20rem] justify-center border-2 p-5">
			<Typography
				variant="h3"
				gutterBottom
				className="text-center text-blue-500 font-semibold">
				ScholarAI
			</Typography>
			
			<Typography variant="h5" className="text-center font-semibold">
				Verify Your Email
			</Typography>
			
			<Typography variant="body1" className="text-center text-gray-600">
				We've sent a 6-digit verification code to your email address.
				Please enter the code below to verify your account.
			</Typography>

			<OTPInput
				onChange={handleOTPChange}
				onComplete={handleOTPComplete}
				disabled={isVerifying}
				error={!!error}
				helperText={error || 'Enter the 6-digit code'}
			/>

			{isVerifying && (
				<Box className="flex justify-center items-center gap-2">
					<CircularProgress size={20} />
					<Typography variant="body2" color="text.secondary">
						Verifying...
					</Typography>
				</Box>
			)}

			<Box className="text-center">
				<Typography variant="body2" color="text.secondary">
					Didn't receive the code?{' '}
					<Button
						variant="text"
						onClick={resendOTP}
						disabled={isVerifying}
						sx={{ textTransform: 'none', p: 0, minWidth: 'auto' }}
					>
						Resend Code
					</Button>
				</Typography>
			</Box>

			<Box className="text-center">
				<Typography variant="body2" color="text.secondary">
					Wrong email?{' '}
					<Link
						href="/sign-up"
						className="text-blue-500 font-semibold hover:underline">
						Go back to sign up
					</Link>
				</Typography>
			</Box>
		</Box>
	);
};

export default VerifyEmailPage;
