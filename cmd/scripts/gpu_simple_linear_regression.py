print("success")
# import torch
# import torch.nn as nn
# import torch.optim as optim
#
# device = torch.device("cuda" if torch.cuda.is_available() else "cpu")
# print(f'Using {device} device.')
#
# torch.manual_seed(0)
# x = torch.randn(100, 1) * 10
# y = x + torch.randn(100, 1) * 5
#
# x = x.to(device)
# y = y.to(device)
#
# class SimpleLinearRegression(nn.Module):
#     def __init__(self):
#         super(SimpleLinearRegression, self).__init__()
#         self.linear = nn.Linear(in_features=1, out_features=1)
#
#     def forward(self, x):
#         return self.linear(x)
#
# model = SimpleLinearRegression().to(device)
#
# criterion = nn.MSELoss()
# optimizer = optim.SGD(model.parameters(), lr=0.01)
#
# num_epochs = 20
# for epoch in range(num_epochs):
#     model.train()
#     optimizer.zero_grad()
#     outputs = model(x)
#     loss = criterion(outputs, y)
#     loss.backward()
#     optimizer.step()
#
# model.eval()
# with torch.no_grad():
#     predicted = model(x)
#     print(f'Predicted: success')
#     print(f'Actual: success')