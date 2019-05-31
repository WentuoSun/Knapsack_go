   1: import java.util.Stack;
   2:  
   3: class HeapNode {
   4:     double upbound; // 结点的价值上界
   5:     double value; // 结点所对应的价值
   6:     double weight; // 结点所相应的重量
   7:  
   8:     int level; // 活节点在子集树中所处的层序号
   9:  
  10:     public HeapNode() {
  11:     }
  12: }
  13:  
  14: // 分支限界法实现01背包问题
  15: public class BB_Knapsack01 {
  16:  
  17:     int[] weight;
  18:     int[] value;
  19:     int max; // 背包的最大承重量
  20:  
  21:     int n;
  22:  
  23:     double c_weight; // 当前背包重量
  24:     double c_value; // 当前背包价值
  25:  
  26:     double bestv; // 最优的背包价值
  27:  
  28:     Stack<HeapNode> heap;
  29:  
  30:     public BB_Knapsack01() {        
  31:         weight = new int[] { 15, 16, 15, 0 };
  32:         value = new int[] { 25, 45, 25, 0 };
  33:         max = 30;
  34:  
  35:         n = weight.length - 1;
  36:  
  37:         c_weight = 0;
  38:         c_value = 0;
  39:         bestv = 0;
  40:  
  41:         heap = new Stack<HeapNode>();
  42:     }
  43:  
  44:     // 求子树的最大上界
  45:     private double maxBound(int t) {
  46:         double left = max - c_weight;
  47:         double bound = c_value;
  48:         // 剩余容量和价值上界
  49:         while (t < n && weight[t] <= left) {
  50:             left -= weight[t];
  51:             bound += value[t];
  52:             t++;
  53:         }
  54:         if (t < n)
  55:             bound += (value[t] / weight[t]) * left; // 装填剩余容量装满背包        
  56:         return bound;
  57:     }
  58:  
  59:     // 将一个新的活结点插入到子集树和最大堆heap中
  60:     private void addLiveNode(double upper, double cvalue, double cweight,
  61:             int level) {
  62:         HeapNode node = new HeapNode();
  63:         node.upbound = upper;
  64:         node.value = cvalue;
  65:         node.weight = cweight;
  66:         node.level = level;
  67:         if (level <= n)
  68:             heap.push(node);
  69:     }
  70:  
  71:     // 利用分支限界法，返回最大价值bestv
  72:     private double knapsack() {
  73:         int i = 0;
  74:         double upbound = maxBound(i);
  75:         // 调用maxBound求出价值上界，bestv为最优值
  76:         while (true) // 非叶子结点        
  77:         {
  78:             double wt = c_weight + weight[i];
  79:             if (wt <= max)// 左儿子结点为可行结点            
  80:             {
  81:                 if (c_value + value[i] > bestv)
  82:                     bestv = c_value + value[i];
  83:                 addLiveNode(upbound, c_value + value[i], c_weight + weight[i],
  84:                         i + 1);
  85:             }
  86:             upbound = maxBound(i + 1);
  87:             if (upbound >= bestv) // 右子树可能含最优解
  88:                 addLiveNode(upbound, c_value, c_weight, i + 1);
  89:             if (heap.empty())
  90:                 return bestv;
  91:             HeapNode node = heap.peek();
  92:             // 取下一扩展结点
  93:             heap.pop();
  94:             //System.out.println(node.value + " ");
  95:             c_weight = node.weight;
  96:             c_value = node.value;
  97:             upbound = node.upbound;
  98:             i = node.level;
  99:         }
 100:     }
 101:  
 102:     public static void main(String[] args) {
 103:         // TODO Auto-generated method stub
 104:         BB_Knapsack01 knap = new BB_Knapsack01();
 105:         double opt_value = knap.knapsack();
 106:         System.out.println(opt_value);
 107:     }
 108: }