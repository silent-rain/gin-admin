<template>
  <!-- Element Icon -->
  <el-tab-pane label="Element UI" name="element">
    <el-row>
      <el-col :span="4" v-for="(item, _) in getOptions">
        <el-dropdown-item :command="item.value">
          <el-icon class="el-svg-icon">
            <component :is="ElSvg[item.value]" />
          </el-icon>
        </el-dropdown-item>
      </el-col>
    </el-row>
  </el-tab-pane>
</template>

<script setup lang="ts">
import * as ElSvg from '@element-plus/icons-vue';

const props = withDefaults(
  defineProps<{
    query: string;
  }>(),
  {},
);

const options = [
  // System
  { label: 'Plus', value: 'Plus' },
  { label: 'Minus', value: 'Minus' },
  { label: 'CirclePlus', value: 'CirclePlus' },
  { label: 'Search', value: 'Search' },
  { label: 'Female', value: 'Female' },
  { label: 'Male', value: 'Male' },
  { label: 'Aim', value: 'Aim' },
  { label: 'House', value: 'House' },
  { label: 'FullScreen', value: 'FullScreen' },
  { label: 'Loading', value: 'Loading' },
  { label: 'Link', value: 'Link' },
  { label: 'Service', value: 'Service' },
  { label: 'Pointer', value: 'Pointer' },
  { label: 'Star', value: 'Star' },
  { label: 'Notification', value: 'Notification' },
  { label: 'Connection', value: 'Connection' },
  { label: 'ChatDotRound', value: 'ChatDotRound' },
  { label: 'Setting', value: 'Setting' },
  { label: 'Clock', value: 'Clock' },
  { label: 'Position', value: 'Position' },
  { label: 'Discount', value: 'Discount' },
  { label: 'Odometer', value: 'Odometer' },
  { label: 'ChatSquare', value: 'ChatSquare' },
  { label: 'ChatRound', value: 'ChatRound' },
  { label: 'ChatLineRound', value: 'ChatLineRound' },
  { label: 'ChatLineSquare', value: 'ChatLineSquare' },
  { label: 'ChatDotSquare', value: 'ChatDotSquare' },
  { label: 'View', value: 'View' },
  { label: 'Hide', value: 'Hide' },
  { label: 'Unlock', value: 'Unlock' },
  { label: 'Lock', value: 'Lock' },
  { label: 'RefreshRight', value: 'RefreshRight' },
  { label: 'RefreshLeft', value: 'RefreshLeft' },
  { label: 'Refresh', value: 'Refresh' },
  { label: 'Bell', value: 'Bell' },
  { label: 'MuteNotification', value: 'MuteNotification' },
  { label: 'User', value: 'User' },
  { label: 'Check', value: 'Check' },
  { label: 'CircleCheck', value: 'CircleCheck' },
  { label: 'Warning', value: 'Warning' },
  { label: 'CircleClose', value: 'CircleClose' },
  { label: 'Close', value: 'Close' },
  { label: 'PieChart', value: 'PieChart' },
  { label: 'More', value: 'More' },
  { label: 'Compass', value: 'Compass' },
  { label: 'Filter', value: 'Filter' },
  { label: 'Switch', value: 'Switch' },
  { label: 'Select', value: 'Select' },
  { label: 'SemiSelect', value: 'SemiSelect' },
  { label: 'CloseBold', value: 'CloseBold' },
  { label: 'EditPen', value: 'EditPen' },
  { label: 'Edit', value: 'Edit' },
  { label: 'Message', value: 'Message' },
  { label: 'MessageBox', value: 'MessageBox' },
  { label: 'TurnOff', value: 'TurnOff' },
  { label: 'Finished', value: 'Finished' },
  { label: 'Delete', value: 'Delete' },
  { label: 'Crop', value: 'Crop' },
  { label: 'SwitchButton', value: 'SwitchButton' },
  { label: 'Operation', value: 'Operation' },
  { label: 'Open', value: 'Open' },
  { label: 'Remove', value: 'Remove' },
  { label: 'ZoomOut', value: 'ZoomOut' },
  { label: 'ZoomIn', value: 'ZoomIn' },
  { label: 'InfoFilled', value: 'InfoFilled' },
  { label: 'CircleCheckFilled', value: 'CircleCheckFilled' },
  { label: 'SuccessFilled', value: 'SuccessFilled' },
  { label: 'WarningFilled', value: 'WarningFilled' },
  { label: 'CircleCloseFilled', value: 'CircleCloseFilled' },
  { label: 'QuestionFilled', value: 'QuestionFilled' },
  { label: 'WarnTriangleFilled', value: 'WarnTriangleFilled' },
  { label: 'UserFilled', value: 'UserFilled' },
  { label: 'MoreFilled', value: 'MoreFilled' },
  { label: 'Tools', value: 'Tools' },
  { label: 'HomeFilled', value: 'HomeFilled' },
  { label: 'Menu', value: 'Menu' },
  { label: 'UploadFilled', value: 'UploadFilled' },
  { label: 'Avatar', value: 'Avatar' },
  { label: 'HelpFilled', value: 'HelpFilled' },
  { label: 'Share', value: 'Share' },
  { label: 'StarFilled', value: 'StarFilled' },
  { label: 'Comment', value: 'Comment' },
  { label: 'Histogram', value: 'Histogram' },
  { label: 'Grid', value: 'Grid' },
  { label: 'Promotion', value: 'Promotion' },
  { label: 'DeleteFilled', value: 'DeleteFilled' },
  { label: 'RemoveFilled', value: 'RemoveFilled' },
  { label: 'CirclePlusFilled', value: 'CirclePlusFilled' },
  // Arrow
  { label: 'ArrowLeft', value: 'ArrowLeft' },
  { label: 'ArrowUp', value: 'ArrowUp' },
  { label: 'ArrowRight', value: 'ArrowRight' },
  { label: 'ArrowDown', value: 'ArrowDown' },
  { label: 'ArrowLeftBold', value: 'ArrowLeftBold' },
  { label: 'ArrowUpBold', value: 'ArrowUpBold' },
  { label: 'ArrowRightBold', value: 'ArrowRightBold' },
  { label: 'ArrowDownBold', value: 'ArrowDownBold' },
  { label: 'DArrowRight', value: 'DArrowRight' },
  { label: 'DArrowLeft', value: 'DArrowLeft' },
  { label: 'Download', value: 'Download' },
  { label: 'Upload', value: 'Upload' },
  { label: 'Top', value: 'Top' },
  { label: 'Bottom', value: 'Bottom' },
  { label: 'Back', value: 'Back' },
  { label: 'Right', value: 'Right' },
  { label: 'TopRight', value: 'TopRight' },
  { label: 'TopLeft', value: 'TopLeft' },
  { label: 'BottomRight', value: 'BottomRight' },
  { label: 'BottomLeft', value: 'BottomLeft' },
  { label: 'Sort', value: 'Sort' },
  { label: 'SortUp', value: 'SortUp' },
  { label: 'SortDown', value: 'SortDown' },
  { label: 'Rank', value: 'Rank' },
  { label: 'CaretLeft', value: 'CaretLeft' },
  { label: 'CaretTop', value: 'CaretTop' },
  { label: 'CaretRight', value: 'CaretRight' },
  { label: 'CaretBottom', value: 'CaretBottom' },
  { label: 'DCaret', value: 'DCaret' },
  { label: 'Expand', value: 'Expand' },
  { label: 'Fold', value: 'Fold' },
  // Document
  { label: 'DocumentAdd', value: 'DocumentAdd' },
  { label: 'Document', value: 'Document' },
  { label: 'Notebook', value: 'Notebook' },
  { label: 'Tickets', value: 'Tickets' },
  { label: 'Memo', value: 'Memo' },
  { label: 'Collection', value: 'Collection' },
  { label: 'Postcard', value: 'Postcard' },
  { label: 'ScaleToOriginal', value: 'ScaleToOriginal' },
  { label: 'SetUp', value: 'SetUp' },
  { label: 'DocumentDelete', value: 'DocumentDelete' },
  { label: 'DocumentChecked', value: 'DocumentChecked' },
  { label: 'DataBoard', value: 'DataBoard' },
  { label: 'DataAnalysis', value: 'DataAnalysis' },
  { label: 'CopyDocument', value: 'CopyDocument' },
  { label: 'FolderChecked', value: 'FolderChecked' },
  { label: 'Files', value: 'Files' },
  { label: 'Folder', value: 'Folder' },
  { label: 'FolderDelete', value: 'FolderDelete' },
  { label: 'FolderRemove', value: 'FolderRemove' },
  { label: 'FolderOpened', value: 'FolderOpened' },
  { label: 'DocumentCopy', value: 'DocumentCopy' },
  { label: 'DocumentRemove', value: 'DocumentRemove' },
  { label: 'FolderAdd', value: 'FolderAdd' },
  { label: 'FirstAidKit', value: 'FirstAidKit' },
  { label: 'Reading', value: 'Reading' },
  { label: 'DataLine', value: 'DataLine' },
  { label: 'Management', value: 'Management' },
  { label: 'Checked', value: 'Checked' },
  { label: 'Ticket', value: 'Ticket' },
  { label: 'Failed', value: 'Failed' },
  { label: 'TrendCharts', value: 'TrendCharts' },
  { label: 'List', value: 'List' },
  // Media
  { label: 'Microphone', value: 'Microphone' },
  { label: 'Mute', value: 'Mute' },
  { label: 'Mic', value: 'Mic' },
  { label: 'VideoPause', value: 'VideoPause' },
  { label: 'VideoCamera', value: 'VideoCamera' },
  { label: 'VideoPlay', value: 'VideoPlay' },
  { label: 'Headset', value: 'Headset' },
  { label: 'Monitor', value: 'Monitor' },
  { label: 'Film', value: 'Film' },
  { label: 'Camera', value: 'Camera' },
  { label: 'Picture', value: 'Picture' },
  { label: 'PictureRounded', value: 'PictureRounded' },
  { label: 'Iphone', value: 'Iphone' },
  { label: 'Cellphone', value: 'Cellphone' },
  { label: 'VideoCameraFilled', value: 'VideoCameraFilled' },
  { label: 'PictureFilled', value: 'PictureFilled' },
  { label: 'Platform', value: 'Platform' },
  { label: 'CameraFilled', value: 'CameraFilled' },
  { label: 'BellFilled', value: 'BellFilled' },
  // Traffic
  { label: 'Location', value: 'Location' },
  { label: 'LocationInformation', value: 'LocationInformation' },
  { label: 'DeleteLocation', value: 'DeleteLocation' },
  { label: 'Coordinate', value: 'Coordinate' },
  { label: 'Bicycle', value: 'Bicycle' },
  { label: 'OfficeBuilding', value: 'OfficeBuilding' },
  { label: 'School', value: 'School' },
  { label: 'Guide', value: 'Guide' },
  { label: 'AddLocation', value: 'AddLocation' },
  { label: 'MapLocation', value: 'MapLocation' },
  { label: 'Place', value: 'Place' },
  { label: 'LocationFilled', value: 'LocationFilled' },
  { label: 'Van', value: 'Van' },
  // Food
  { label: 'Watermelon', value: 'Watermelon' },
  { label: 'Pear', value: 'Pear' },
  { label: 'NoSmoking', value: 'NoSmoking' },
  { label: 'Smoking', value: 'Smoking' },
  { label: 'Mug', value: 'Mug' },
  { label: 'GobletSquareFull', value: 'GobletSquareFull' },
  { label: 'GobletFull', value: 'GobletFull' },
  { label: 'KnifeFork', value: 'KnifeFork' },
  { label: 'Sugar', value: 'Sugar' },
  { label: 'Bowl', value: 'Bowl' },
  { label: 'MilkTea', value: 'MilkTea' },
  { label: 'Lollipop', value: 'Lollipop' },
  { label: 'Coffee', value: 'Coffee' },
  { label: 'Chicken', value: 'Chicken' },
  { label: 'Dish', value: 'Dish' },
  { label: 'IceTea', value: 'IceTea' },
  { label: 'ColdDrink', value: 'ColdDrink' },
  { label: 'CoffeeCup', value: 'CoffeeCup' },
  { label: 'DishDot', value: 'DishDot' },
  { label: 'IceDrink', value: 'IceDrink' },
  { label: 'IceCream', value: 'IceCream' },
  { label: 'Dessert', value: 'Dessert' },
  { label: 'IceCreamSquare', value: 'IceCreamSquare' },
  { label: 'ForkSpoon', value: 'ForkSpoon' },
  { label: 'IceCreamRound', value: 'IceCreamRound' },
  { label: 'Food', value: 'Food' },
  { label: 'HotWater', value: 'HotWater' },
  { label: 'Grape', value: 'Grape' },
  { label: 'Fries', value: 'Fries' },
  { label: 'Apple', value: 'Apple' },
  { label: 'Burger', value: 'Burger' },
  { label: 'Goblet', value: 'Goblet' },
  { label: 'GobletSquare', value: 'GobletSquare' },
  { label: 'Orange', value: 'Orange' },
  { label: 'Cherry', value: 'Cherry' },
  // Items
  { label: 'Printer', value: 'Printer' },
  { label: 'Calendar', value: 'Calendar' },
  { label: 'CreditCard', value: 'CreditCard' },
  { label: 'Box', value: 'Box' },
  { label: 'Money', value: 'Money' },
  { label: 'Refrigerator', value: 'Refrigerator' },
  { label: 'Cpu', value: 'Cpu' },
  { label: 'Football', value: 'Football' },
  { label: 'Brush', value: 'Brush' },
  { label: 'Suitcase', value: 'Suitcase' },
  { label: 'SuitcaseLine', value: 'SuitcaseLine' },
  { label: 'Umbrella', value: 'Umbrella' },
  { label: 'AlarmClock', value: 'AlarmClock' },
  { label: 'Medal', value: 'Medal' },
  { label: 'GoldMedal', value: 'GoldMedal' },
  { label: 'Present', value: 'Present' },
  { label: 'Mouse', value: 'Mouse' },
  { label: 'Watch', value: 'Watch' },
  { label: 'QuartzWatch', value: 'QuartzWatch' },
  { label: 'Magnet', value: 'Magnet' },
  { label: 'Help', value: 'Help' },
  { label: 'Soccer', value: 'Soccer' },
  { label: 'ToiletPaper', value: 'ToiletPaper' },
  { label: 'ReadingLamp', value: 'ReadingLamp' },
  { label: 'Paperclip', value: 'Paperclip' },
  { label: 'MagicStick', value: 'MagicStick' },
  { label: 'Basketball', value: 'Basketball' },
  { label: 'Baseball', value: 'Baseball' },
  { label: 'Coin', value: 'Coin' },
  { label: 'Goods', value: 'Goods' },
  { label: 'Sell', value: 'Sell' },
  { label: 'SoldOut', value: 'SoldOut' },
  { label: 'Key', value: 'Key' },
  { label: 'ShoppingCart', value: 'ShoppingCart' },
  { label: 'ShoppingCartFull', value: 'ShoppingCartFull' },
  { label: 'ShoppingTrolley', value: 'ShoppingTrolley' },
  { label: 'Phone', value: 'Phone' },
  { label: 'Scissor', value: 'Scissor' },
  { label: 'Handbag', value: 'Handbag' },
  { label: 'ShoppingBag', value: 'ShoppingBag' },
  { label: 'Trophy', value: 'Trophy' },
  { label: 'TrophyBase', value: 'TrophyBase' },
  { label: 'Stopwatch', value: 'Stopwatch' },
  { label: 'Timer', value: 'Timer' },
  { label: 'CollectionTag', value: 'CollectionTag' },
  { label: 'TakeawayBox', value: 'TakeawayBox' },
  { label: 'PriceTag', value: 'PriceTag' },
  { label: 'Wallet', value: 'Wallet' },
  { label: 'Opportunity', value: 'Opportunity' },
  { label: 'PhoneFilled', value: 'PhoneFilled' },
  { label: 'WalletFilled', value: 'WalletFilled' },
  { label: 'GoodsFilled', value: 'GoodsFilled' },
  { label: 'Flag', value: 'Flag' },
  { label: 'BrushFilled', value: 'BrushFilled' },
  { label: 'Briefcase', value: 'Briefcase' },
  { label: 'Stamp', value: 'Stamp' },
  // Weather
  { label: 'Sunrise', value: 'Sunrise' },
  { label: 'Sunny', value: 'Sunny' },
  { label: 'Ship', value: 'Ship' },
  { label: 'MostlyCloudy', value: 'MostlyCloudy' },
  { label: 'PartlyCloudy', value: 'PartlyCloudy' },
  { label: 'Sunset', value: 'Sunset' },
  { label: 'Drizzling', value: 'Drizzling' },
  { label: 'Pouring', value: 'Pouring' },
  { label: 'Cloudy', value: 'Cloudy' },
  { label: 'Moon', value: 'Moon' },
  { label: 'MoonNight', value: 'MoonNight' },
  { label: 'Lightning', value: 'Lightning' },
  // Other
  { label: 'ChromeFilled', value: 'ChromeFilled' },
  { label: 'Eleme', value: 'Eleme' },
  { label: 'ElemeFilled', value: 'ElemeFilled' },
  { label: 'ElementPlus', value: 'ElementPlus' },
  { label: 'Shop', value: 'Shop' },
  { label: 'SwitchFilled', value: 'SwitchFilled' },
  { label: 'WindPower', value: 'WindPower' },
];

const getOptions = computed(() => {
  if (props.query === '') {
    return options;
  }
  const re = new RegExp(`${props.query}`, 'i');
  return options.filter((v) => re.test(v.value));
});
</script>

<style scoped lang="scss">
.el-svg-icon {
  height: 45px;
  width: 45px;
  text-align: center;
  font-size: 25px;
  &:hover {
    transform: scale(1.2);
  }
}
</style>
