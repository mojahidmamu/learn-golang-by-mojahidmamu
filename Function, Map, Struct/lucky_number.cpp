#include <bits/stdc++.h>
using namespace std;

vector<int> luckyNumbers(int n)
{
    vector<int> nums;
    for (int i = 1; i <= n * 10; i += 2)
        nums.push_back(i);

    for (int i = 1; i < nums.size(); i++)
    {
        int step = nums[i];
        for (int j = step - 1; j < nums.size(); j += step - 1)
            nums.erase(nums.begin() + j);
    }
    return nums;
}

int main()
{
    int n;
    cin >> n;
    vector<int> lucky = luckyNumbers(n);
    for (int i = 0; i < n; i++)
        cout << lucky[i] << " ";
}
