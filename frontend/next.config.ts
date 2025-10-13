import type { NextConfig } from 'next';

const nextConfig: NextConfig = {
	/* config options here */
	env: {
		NEXT_PUBLIC_API_URL: 'localhost:8080/api/v1',
	},
};

export default nextConfig;
