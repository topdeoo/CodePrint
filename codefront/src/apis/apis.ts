import service from './server/service';
export const teamLogin=(info:any) => service.post(
  '/login', info
);
export const codePrint = (code:any) => service.post(
  '/print', code
)