import {
    BuildingLibraryIcon,
    FolderIcon,
    HomeIcon,
    InboxIcon,
    UserGroupIcon,
    UsersIcon
} from '@heroicons/vue/24/outline';

export const navigation = [
    {name: 'Dashboard', href: '/', current: '', icon: HomeIcon},
    {name: 'Users', href: '/users', current: 'users', icon: UsersIcon},
];

export const notifications = [
    {title: 'Assessment due', body: 'ADAM assessment due in 5 days'},
    {title: 'New assignment added', body: 'Sight words assignment added'},
];

export const userNavigation = [
    {name: 'Profile', href: '/profile', current: 'profile'},
    {name: 'Logout', href: '/logout', current: 'logout'},
]