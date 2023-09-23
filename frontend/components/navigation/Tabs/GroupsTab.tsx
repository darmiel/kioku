import { Group } from "../../../types/Group";
import DeckOverview from "../../deck/DeckOverview";

interface GroupsTabProps {
	/**
	 * groups
	 */
	groups: Group[];
	/**
	 * Additional classes
	 */
	className?: string;
}

/**
 * UI component for the GroupsTab
 */
export const GroupsTab = ({ groups, className = "" }: GroupsTabProps) => {
	return (
		<div className={`${className}`}>
			{groups
				?.filter((group: Group) => !group.isDefault)
				.map((group: Group) => {
					return (
						<DeckOverview
							key={group.groupID}
							group={group}
						></DeckOverview>
					);
				})}
			<DeckOverview />
		</div>
	);
};
