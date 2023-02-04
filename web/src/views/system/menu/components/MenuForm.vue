<template>
  <el-dialog
    :model-value="props.visible"
    :title="props.type === 'add' ? '添加菜单' : '编辑菜单'"
    :width="props.width"
    :before-close="handleClose"
  >
    <el-form
      ref="ruleFormRef"
      :rules="rules"
      :model="props.data"
      label-width="100px"
      style="width: 100%"
    >
      <el-row>
        <el-col :span="12">
          <el-form-item label="上级菜单" prop="parent_id">
            <el-tree-select
              v-model="props.data.parent_id"
              :data="menuOptions"
              node-key="id"
              :props="{
                children: 'children',
                label: 'title',
              }"
              :render-after-expand="false"
              filterable
              accordion
              :check-strictly="true"
              style="width: 100%"
            >
              <template #default="{ node, _data }">
                <span class="custom-tree-node">
                  <span>{{ node.label }}</span>
                </span>
              </template>
            </el-tree-select>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="菜单类型" prop="menu_type">
            <el-radio-group v-model="props.data.menu_type">
              <el-radio :label="0">菜单</el-radio>
              <el-radio :label="1">按钮</el-radio>
            </el-radio-group>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="菜单名称" prop="title">
            <el-input v-model="props.data.title" placeholder="请输入菜单名称" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="打开方式" prop="open_type">
            <el-radio-group
              v-model="props.data.open_type"
              :disabled="props.data.menu_type === MenuType.Button"
            >
              <el-radio :label="0">组件</el-radio>
              <el-radio :label="1">内链</el-radio>
              <el-radio :label="2">外链</el-radio>
            </el-radio-group>
          </el-form-item>
        </el-col>

        <el-divider />

        <el-col :span="12">
          <el-form-item label="菜单图标" prop="icon">
            <el-input
              v-model="props.data.icon"
              :disabled="props.data.menu_type === MenuType.Button"
              placeholder="请选择菜单图标"
            />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="权限标识" prop="permission">
            <el-input
              v-model="props.data.permission"
              :disabled="props.data.menu_type !== MenuType.Button"
              placeholder="请输入权限标识"
            />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="路由地址" prop="path">
            <el-input
              v-model="props.data.path"
              :disabled="
                props.data.menu_type === MenuType.Button ||
                props.data.open_type === OpenType.OuterLink
              "
              placeholder="请输入路由地址"
            />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="排序号" prop="sort">
            <el-input-number v-model="props.data.sort" :min="1" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item
            v-if="props.data.open_type === OpenType.Component"
            label="组件路径"
            prop="component"
          >
            <el-input
              v-model="props.data.component"
              :disabled="props.data.menu_type === MenuType.Button"
              placeholder="请输入组件路径"
            />
          </el-form-item>
          <el-form-item
            v-if="props.data.open_type === OpenType.Link"
            label="内链地址"
            prop="link"
          >
            <template #label>
              <div>
                <el-tooltip
                  content="需要以`http://`、`https://`、`//`开头"
                  placement="top"
                >
                  <el-icon style="margin-right: 2px">
                    <QuestionFilled
                  /></el-icon>
                </el-tooltip>
                <span>内链地址</span>
              </div>
            </template>
            <el-input
              v-model="props.data.link"
              :disabled="props.data.menu_type === MenuType.Button"
              placeholder="请输入内链地址"
            />
          </el-form-item>
          <el-form-item
            v-if="props.data.open_type === OpenType.OuterLink"
            label="外链地址"
            prop="link"
          >
            <template #label>
              <div>
                <el-tooltip
                  content="需要以`http://`、`https://`、`//`开头"
                  placement="top"
                >
                  <el-icon style="margin-right: 2px">
                    <QuestionFilled
                  /></el-icon>
                </el-tooltip>
                <span>外链地址</span>
              </div>
            </template>
            <el-input
              v-model="props.data.link"
              :disabled="props.data.menu_type === MenuType.Button"
              placeholder="请输入外链地址"
            />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="是否可见" prop="hide">
            <el-radio-group
              v-model="props.data.hide"
              :disabled="props.data.menu_type === MenuType.Button"
            >
              <el-radio :label="1">显示</el-radio>
              <el-radio :label="0"
                >隐藏
                <el-tooltip
                  content="选择不可见只注册路由不显示在侧边栏"
                  placement="top"
                >
                  <el-icon style="margin-left: 1px">
                    <QuestionFilled
                  /></el-icon>
                </el-tooltip>
              </el-radio>
            </el-radio-group>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="启用状态" prop="status">
            <el-radio-group v-model="props.data.status">
              <el-radio :label="1">启用</el-radio>
              <el-radio :label="0">停用</el-radio>
            </el-radio-group>
          </el-form-item>
        </el-col>
        <el-col :span="24">
          <el-form-item label="备注" prop="note">
            <el-input
              v-model="props.data.note"
              type="textarea"
              placeholder="请输入备注"
            />
          </el-form-item>
        </el-col>
        <el-col :span="12"></el-col>
      </el-row>
    </el-form>

    <template #footer>
      <span class="dialog-footer">
        <el-button @click="handleCancel">取消</el-button>
        <el-button type="primary" @click="submitForm(ruleFormRef)">
          提交
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ElMessage, FormInstance, FormRules } from 'element-plus';
import { QuestionFilled } from '@element-plus/icons-vue';
import { updateMenu, addMenu, getAllMenu } from '@/api/system/menu';
import { Menu, MenuListRsp } from '~/api/system/menu';
import { MenuType, OpenType } from '@/constant/system/menu';

const emit = defineEmits(['update:data', 'update:visible', 'refresh']);

const props = withDefaults(
  defineProps<{
    data: Menu;
    visible: boolean;
    type: string; // add/edit
    width?: string;
  }>(),
  {
    width: '100%',
  },
);

const ruleFormRef = ref<FormInstance>();
const rules = reactive<FormRules>({
  title: [
    { required: true, message: '请输入菜单名称', trigger: 'blur' },
    { min: 2, message: '至少输入两个字符', trigger: 'blur' },
  ],
});
// 菜单列表
const menuOptions = ref<Menu[]>([]);

onBeforeMount(() => {
  fetchAllMenu();
});

// 获取所有菜单列表
const fetchAllMenu = async () => {
  try {
    const resp = (await getAllMenu()).data as MenuListRsp;
    menuOptions.value = resp.data_list.filter((v: any) => {
      // 过滤自身选择, 防止自依赖
      if (v.id !== props.data.id) {
        return true;
      }
    });
  } catch (error) {
    console.log(error);
  }
};

// 关闭
const handleClose = () => {
  emit('update:visible', false);
  emit('update:data', {});
};

// 取消
const handleCancel = () => {
  emit('update:visible', false);
  emit('update:data', {});
};
// 提交
const submitForm = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  await formEl.validate(async (valid, fields) => {
    if (!valid) {
      console.log('error submit!', fields);
      return;
    }

    try {
      if (props.type === 'add') {
        await addMenu(props.data);
      } else {
        await updateMenu(props.data);
      }
      fetchAllMenu();
      emit('update:visible', false);
      emit('update:data', {});
      emit('refresh');
      ElMessage.success('操作成功');
    } catch (error) {
      console.log(error);
    }
  });
};
</script>

<style scoped lang="scss">
.el-divider--horizontal {
  margin: 10px, 0;
}
</style>
