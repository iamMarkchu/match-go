# 篮球计分技术统计软件

## 需求点
1. 新建一场比赛
2. 配置比赛
    - 比赛双方队名
    - 比赛倒计时(可以设置节数，每节分钟数)，支持暂停
    - 比赛双方队员录入
3. 比赛记录
    - 得分记录 (1分，2分，3分)
    - 犯规记录, 犯规队员记录
    - 暂停比赛，开始比赛
    - 记录暂停时间
    - 技术统计
4. 赛后统计
5. 修改比赛配置



match // 比赛表
- id
- name  // 比赛名称
- pic   // 宣传图
- desc  // 比赛简介
- status
- utime
- ctime
- opr

match_config // 比赛配置表
- id
- mid
- player_config
  {
  "match_a": "",
  "match_b": "",
  "match_a_players": [
  {
  "number": 99,
  "name": "ck",
  "height": "179",
  "weight": "85",
  "desc": "技术特点，球员简介，可为空",
  "avatar": "头像",
  "is_first": 1, // 是否首发
  }
  ]
  }
- status
- utime
- ctime
- opr

match_log // 比赛流水表
- id
- action  // 动作名称 1_SHOOT_ONE: 罚球得分 1_SHOOT_TWO: 两分得分 1_SHOOT_THREE: 三分得分, 2_FOUL: 犯规, 3_REBOUND:篮板 3_STEAL:抢断 3_SHOOT_ONE:罚球投篮 3_Shoot_TWO:两分投篮 3_Shoot_THREE:三分投篮
- action_type // 动作类型, 1:得分 2:犯规 3.技术统计
- key // 队员
- value // 数值, 默认1
- status // 是否有效
- time_str // 第几节，第几分钟
- ctime  // 发生时间
- opr  // 记录人
