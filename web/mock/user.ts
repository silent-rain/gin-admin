import Mock from "mockjs";

export default [
  {
    url: "/basis-func/user/getUserInfo",
    method: "post",
    response: () => {
      return {
        msg: "操作成功!",
        flag: true,
        code: 20000,
        data: {
          userInfo: {
            deleted: 0,
            headImgUrl:
              "https://github.jzfai.top/gofast-image//group1/default/20221108/19/12/4/logo_W2xEX_2x.jpg",
            phone: "13302254696",
            createTime: 1666410370000,
            roleId: "[13]",
            name: "panda",
            updateTime: 1669861280000,
            id: 5,
          },
          codes: [16, 9, 10, 11, 12, 13, 15],
          menuList: [
            {
              redirect: "",
              code: "RbacTest",
              hidden: 0,
              icon: "",
              updateTime: "2022-12-29 15:53:01",
              parentNode: 1,
              title: "RBACMenu",
              parentId: 0,
              plateFormId: 2,
              path: "rbac-test",
              component: "Layout",
              deleted: 0,
              createTime: "2022-10-15 17:41:23",
              children: [
                {
                  redirect: "",
                  code: "TestMenu",
                  hidden: 0,
                  icon: "",
                  updateTime: "2022-10-25 14:17:28",
                  parentNode: 1,
                  sort: 10,
                  title: "路由权限测试",
                  parentId: 9,
                  plateFormId: 2,
                  path: "test-table-query",
                  component: "rbac-test/TestMenu.vue",
                  deleted: 0,
                  createTime: "2022-10-15 17:58:09",
                  children: [],
                  intro: "",
                  name: "页面测试",
                  id: 10,
                  elSvgIcon: "",
                  category: 1,
                  alwaysShow: 0,
                },
                {
                  redirect: "",
                  code: "TestAddEdit",
                  hidden: 1,
                  icon: "",
                  updateTime: "2022-10-25 14:00:45",
                  sort: 10,
                  title: "页面测试-新增编辑",
                  parentId: 9,
                  plateFormId: 2,
                  path: "test-add-edit",
                  component: "rbac-test/TestAddEdit.vue",
                  deleted: 0,
                  createTime: "2022-10-15 18:00:26",
                  intro: "",
                  name: "页面测试-新增编辑",
                  id: 11,
                  elSvgIcon: "",
                  category: 1,
                  alwaysShow: 0,
                },
                {
                  redirect: "",
                  code: "TestButton",
                  hidden: 0,
                  icon: "",
                  updateTime: "2022-10-28 16:47:29",
                  parentNode: 1,
                  sort: 10,
                  title: "按钮测试",
                  parentId: 9,
                  plateFormId: 2,
                  path: "test-button",
                  component: "rbac-test/TestButton.vue",
                  deleted: 0,
                  createTime: "2022-10-23 17:28:10",
                  children: [
                    {
                      redirect: "",
                      code: "TestButton:Edit",
                      hidden: 0,
                      icon: "",
                      updateTime: "2022-10-25 14:18:31",
                      parentNode: 0,
                      sort: 0,
                      title: "",
                      parentId: 13,
                      plateFormId: 2,
                      path: "",
                      component: "",
                      deleted: 0,
                      createTime: "2022-10-23 17:36:21",
                      intro: "",
                      extra: "",
                      name: "按钮测试-编辑按钮",
                      id: 15,
                      elSvgIcon: "",
                      category: 3,
                      alwaysShow: 0,
                    },
                    {
                      redirect: "",
                      code: "TestButton:AddBtn",
                      hidden: 0,
                      icon: "",
                      updateTime: "2022-10-25 14:18:38",
                      title: "新增按钮",
                      parentId: 13,
                      plateFormId: 2,
                      path: "",
                      component: "",
                      deleted: 0,
                      createTime: "2022-10-15 18:08:44",
                      intro: "TestTableQuery页面的新增按钮",
                      name: "按钮测试-新增按钮",
                      id: 12,
                      elSvgIcon: "",
                      category: 3,
                      alwaysShow: 1,
                    },
                  ],
                  intro: "",
                  extra: "",
                  name: "按钮测试",
                  id: 13,
                  elSvgIcon: "",
                  category: 1,
                  alwaysShow: 0,
                },
                {
                  redirect: "",
                  code: "TestDetail",
                  hidden: 1,
                  icon: "",
                  updateTime: "2022-10-25 14:00:54",
                  parentNode: 0,
                  sort: 10,
                  title: "页面测试-详情",
                  parentId: 9,
                  plateFormId: 2,
                  path: "detail",
                  component: "rbac-test/TestDetail.vue",
                  deleted: 0,
                  createTime: "2022-10-23 18:23:53",
                  intro: "",
                  extra: "",
                  name: "页面测试-详情",
                  id: 16,
                  elSvgIcon: "",
                  category: 1,
                  alwaysShow: 0,
                },
              ],
              intro: "",
              name: "RBACMenu",
              id: 9,
              elSvgIcon: "Fold",
              category: 1,
              alwaysShow: 1,
            },
          ],
          roles: ["admin"],
          exp: 1673278949,
          iat: 1673019749,
        },
      };
    },
  },
  {
    url: "/basis-func/user/loginValid",
    method: "post",
    response: () => {
      return {
        msg: "操作成功!",
        flag: true,
        code: 20000,
        data: {
          userInfo: {
            deleted: 0,
            headImgUrl:
              "https://github.jzfai.top/gofast-image//group1/default/20221108/19/12/4/logo_W2xEX_2x.jpg",
            phone: "13302254696",
            createTime: "2022-10-22 11:46:10",
            roleId: "[13]",
            name: "panda",
            updateTime: "2022-12-01 10:21:20",
            id: 5,
          },
          jwtToken:
            "eyJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE2NzMwMTk3MDIsImV4cCI6MTY3MzI3ODkwMiwidXNlckluZm8iOnsiaWQiOjUsIm5hbWUiOiJwYW5kYSIsImhlYWRJbWdVcmwiOiJodHRwczovL2dpdGh1Yi5qemZhaS50b3AvZ29mYXN0LWltYWdlLy9ncm91cDEvZGVmYXVsdC8yMDIyMTEwOC8xOS8xMi80L2xvZ29fVzJ4RVhfMnguanBnIiwicGhvbmUiOiIxMzMwMjI1NDY5NiIsInNhbHQiOm51bGwsInBhc3N3b3JkIjpudWxsLCJyb2xlSWQiOiJbMTNdIiwiY3JlYXRlVGltZSI6MTY2NjQxMDM3MDAwMCwiY3JlYXRvciI6bnVsbCwidXBkYXRlVGltZSI6MTY2OTg2MTI4MDAwMCwiZWRpdG9yIjpudWxsLCJkZWxldGVkIjowfX0.Lx6DLarbyFMR9EMWttncueoXbF5nTeyKl6S5bdFvduA",
        },
      };
    },
  },
];
