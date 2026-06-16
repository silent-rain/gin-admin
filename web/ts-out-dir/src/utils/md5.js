import { md5 } from 'js-md5';
import { SECRET } from '@/constant/permission/auth';
export const md5Encode = (v) => {
    return md5(SECRET + v);
};
