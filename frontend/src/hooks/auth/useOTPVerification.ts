/**
 * Custom hook for OTP verification
 */

import { verifyEmail } from '@/lib/auth';
import { useState } from 'react';

export const useOTPVerification = () => {
  const [otp, setOtp] = useState<string>('');
  const [isVerifying, setIsVerifying] = useState(false);
  const [error, setError] = useState<string>('');
  const [isVerified, setIsVerified] = useState(false);

  /**
   * Handle OTP input change
   */
  const handleOTPChange = (value: string) => {
    setOtp(value);
    setError(''); // Clear error when user types
  };

  /**
   * Handle OTP completion (when all 6 digits are entered)
   */
  const handleOTPComplete = async (value: string) => {
    setIsVerifying(true);
    setError('');

    try {
      const result = await verifyEmail(value);
      
      if (result.success) {
        setIsVerified(true);
        console.log('Email verified successfully');
      } else {
        setError(result.message || 'Verification failed');
      }
    } catch (error) {
      console.error('OTP verification error:', error);
      setError('An unexpected error occurred');
    } finally {
      setIsVerifying(false);
    }
  };

  /**
   * Resend OTP
   */
  const resendOTP = async () => {
    // TODO: Implement resend OTP API call
    console.log('Resending OTP...');
  };

  /**
   * Reset OTP state
   */
  const resetOTP = () => {
    setOtp('');
    setError('');
    setIsVerified(false);
    setIsVerifying(false);
  };

  return {
    otp,
    isVerifying,
    error,
    isVerified,
    handleOTPChange,
    handleOTPComplete,
    resendOTP,
    resetOTP,
  };
};
