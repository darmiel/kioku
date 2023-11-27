import {
	BarChart2,
	Copy,
	Home,
	Layers,
	Mail,
	Settings,
	Users,
} from "react-feather";

interface TabHeaderProps {
	/**
	 * Unique identifier
	 */
	id: string;
	/**
	 * Name
	 */
	name: string;
	/**
	 * Style
	 */
	style: keyof typeof getIcon;
	/**
	 * Text that should be displayed as notification
	 */
	notificationBadgeContent?: string;
	/**
	 * Additional classes
	 */
	className?: string;
}

const getIcon = {
	cards: <Copy size={20} />,
	dashboard: <Home size={20} />,
	decks: <Layers size={20} />,
	groups: <Users size={20} />,
	invitations: <Mail size={20} />,
	settings: <Settings size={20} />,
	statistics: <BarChart2 size={20} />,
	user: <Users size={20} />,
} as const;

/**
 * UI component for a TabHeader
 */
export const TabHeader = ({
	name,
	style,
	notificationBadgeContent = "",
	className = "",
	...props
}: TabHeaderProps) => {
	return (
		<div
			className={`flex flex-row items-center space-x-2 ${className}`}
			{...props}
		>
			{getIcon[style]}
			<div>{name}</div>
			{notificationBadgeContent && (
				<div className="relative flex h-full text-sm text-eggshell">
					<div className="absolute inline-flex h-full w-full animate-[ping_1s_ease-out_3] rounded-full bg-kiokuRed opacity-75" />
					<div className="relative flex h-full w-full justify-center rounded-full bg-kiokuRed px-2">
						{notificationBadgeContent}
					</div>
				</div>
			)}
		</div>
	);
};
