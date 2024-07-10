import { defineStore } from 'pinia';
import { Session } from '@/utils/cache/index';

const teamInfo = JSON.parse(window.localStorage.getItem('teamInfo') || '{}');

export const useTeamInfoStore = defineStore('teamInfo', {
  state: () => ({
    teamname: teamInfo.teamname || '',
    location: teamInfo.location || '',
    token: teamInfo.token || '',
  }),
  persist: true,
  actions: {
    setTeamInfo(teamname: string, location: string, token: string) {
      this.teamname = teamname;
      this.token = token;
      this.location = location;
      Session.set('teamInfo', { teamname, token, location });
    },
    getTeamInfo() {
      return {
        teamname: this.teamname,
        location: this.location,
        token: this.token,
      };
    },
    clearTeamInfo() {
      this.teamname = '';
      this.token = '';
      this.location = '';
      Session.clear();
    }
  },
});
