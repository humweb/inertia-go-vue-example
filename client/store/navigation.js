import {BuildingLibraryIcon, HomeIcon, UsersIcon} from '@heroicons/vue/24/outline';

export const navigation = [
    {name: 'Dashboard', href: '/', current: '', icon: HomeIcon},
    {name: 'Users', href: '/users', current: 'users', icon: UsersIcon},
    {name: 'About', href: '/about', current: 'about', icon: BuildingLibraryIcon},
];
