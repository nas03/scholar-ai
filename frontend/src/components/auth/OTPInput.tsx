/**
 * OTP Input Component - 6 digit verification boxes
 */

import { Box, TextField, Typography } from '@mui/material';
import React, { useEffect, useRef, useState } from 'react';

interface OTPInputProps {
	length?: number;
	onComplete: (otp: string) => void;
	onChange?: (otp: string) => void;
	disabled?: boolean;
	error?: boolean;
	helperText?: string;
}

export const OTPInput: React.FC<OTPInputProps> = ({
	length = 6,
	onComplete,
	onChange,
	disabled = false,
	error = false,
	helperText,
}) => {
	const [otp, setOtp] = useState<string[]>(new Array(length).fill(''));
	const [activeIndex, setActiveIndex] = useState<number>(0);
	const inputRefs = useRef<(HTMLDivElement | null)[]>([]);

	useEffect(() => {
		// Focus first input on mount
		if (inputRefs.current[0]) {
			const input = inputRefs.current[0].querySelector('input');
			input?.focus();
		}
	}, []);

	const handleChange = (index: number, value: string) => {
		// Only allow single digit
		if (value.length > 1) {
			value = value.slice(-1);
		}

		// Only allow numbers
		if (!/^\d*$/.test(value)) {
			return;
		}

		const newOtp = [...otp];
		newOtp[index] = value;
		setOtp(newOtp);

		const otpString = newOtp.join('');
		onChange?.(otpString);

		// Move to next input if current is filled
		if (value && index < length - 1) {
			setActiveIndex(index + 1);
			const nextInput = inputRefs.current[index + 1]?.querySelector('input');
			nextInput?.focus();
		}

		// Check if OTP is complete
		if (otpString.length === length && !otpString.includes('')) {
			onComplete(otpString);
		}
	};

	const handleKeyDown = (index: number, e: React.KeyboardEvent) => {
		// Handle backspace
		if (e.key === 'Backspace') {
			if (!otp[index] && index > 0) {
				// Move to previous input if current is empty
				setActiveIndex(index - 1);
				const prevInput = inputRefs.current[index - 1]?.querySelector('input');
				prevInput?.focus();
			} else {
				// Clear current input
				const newOtp = [...otp];
				newOtp[index] = '';
				setOtp(newOtp);
				onChange?.(newOtp.join(''));
			}
		}

		// Handle arrow keys
		if (e.key === 'ArrowLeft' && index > 0) {
			setActiveIndex(index - 1);
			const prevInput = inputRefs.current[index - 1]?.querySelector('input');
			prevInput?.focus();
		}
		if (e.key === 'ArrowRight' && index < length - 1) {
			setActiveIndex(index + 1);
			const nextInput = inputRefs.current[index + 1]?.querySelector('input');
			nextInput?.focus();
		}
	};

	const handlePaste = (e: React.ClipboardEvent) => {
		e.preventDefault();
		const pastedData = e.clipboardData.getData('text').slice(0, length);

		if (/^\d+$/.test(pastedData)) {
			const newOtp = pastedData
				.split('')
				.concat(new Array(length - pastedData.length).fill(''));
			setOtp(newOtp);
			onChange?.(pastedData);

			// Focus the next empty input or the last input
			const nextIndex = Math.min(pastedData.length, length - 1);
			setActiveIndex(nextIndex);
			const nextInput = inputRefs.current[nextIndex]?.querySelector('input');
			nextInput?.focus();

			if (pastedData.length === length) {
				onComplete(pastedData);
			}
		}
	};

	return (
		<Box>
			<Box
				sx={{
					display: 'flex',
					gap: 1,
					justifyContent: 'center',
					alignItems: 'center',
					marginBottom: 2,
				}}>
				{otp.map((digit, index) => (
					<TextField
						key={index}
						ref={(el) => {
							inputRefs.current[index] = el;
						}}
						value={digit}
						onChange={(e) => handleChange(index, e.target.value)}
						onKeyDown={(e) => handleKeyDown(index, e)}
						onPaste={handlePaste}
						onFocus={() => setActiveIndex(index)}
						disabled={disabled}
						error={error}
						inputProps={{
							maxLength: 1,
							style: {
								textAlign: 'center',
								fontSize: '1.5rem',
								fontWeight: 'bold',
								padding: '12px',
								width: '48px',
								height: '48px',
							},
						}}
						sx={{
							'& .MuiOutlinedInput-root': {
								borderRadius: '8px',
								'& fieldset': {
									borderColor: error ? 'error.main' : 'grey.300',
									borderWidth: '2px',
								},
								'&:hover fieldset': {
									borderColor: error ? 'error.main' : 'primary.main',
								},
								'&.Mui-focused fieldset': {
									borderColor: error ? 'error.main' : 'primary.main',
									borderWidth: '2px',
								},
							},
						}}
					/>
				))}
			</Box>
			{helperText && (
				<Typography
					variant="body2"
					color={error ? 'error.main' : 'text.secondary'}
					textAlign="center"
					sx={{ mt: 1 }}>
					{helperText}
				</Typography>
			)}
		</Box>
	);
};
