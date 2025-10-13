import { cn } from '@/lib/utils';
import { NavBarIProps } from '@/types/auth/auth.type';

import Link from 'next/link';

const NavBar: React.FC<NavBarIProps> = ({ className }) => {
	return (
		<div
			className={cn('w-full flex flex-row gap-3 justify-between', className)}>
			<div className="flex flex-row gap-5">
				<button>Home</button>
				<button>Courses</button>
				<button>Schedules</button>
			</div>
			<div className="flex flex-row justify-end gap-5">
				<button>
					<Link href={'/sign-in'}>Sign in</Link>
				</button>
				<button>
					<Link href={'/sign-up'}>Sign up</Link>
				</button>
			</div>
		</div>
	);
};

export default NavBar;
