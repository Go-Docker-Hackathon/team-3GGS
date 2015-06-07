//
//  MateViewController.m
//  mate-ios
//
//  Created by yangwm on 6/6/15.
//  Copyright (c) 2015 xcodecraft. All rights reserved.
//

#import "MateViewController.h"
#import "ASIHTTPRequest.h"

@interface MateViewController ()
- (IBAction)MateListViewController:(id)sender;

@end

@implementation MateViewController

- (id)initWithNibName:(NSString *)nibNameOrNil bundle:(NSBundle *)nibBundleOrNil
{
    self = [super initWithNibName:nibNameOrNil bundle:nibBundleOrNil];
    if (self) {
        // Custom initialization
    }
    return self;
}

- (void)viewDidLoad
{
    [super viewDidLoad];
    // Do any additional setup after loading the view from its nib.
}

- (void)didReceiveMemoryWarning
{
    [super didReceiveMemoryWarning];
    // Dispose of any resources that can be recreated.
}

- (IBAction)MateListViewController:(id)sender {
    
    NSURL *url = [NSURL URLWithString:@"http://3ggs-mate.daoapp.io/User/MateList?phone_number=13811229996"];
    ASIHTTPRequest *request = [ASIHTTPRequest requestWithURL:url];
    [request startSynchronous];
    NSError *error = [request error];
    if (!error) {
        NSString *response = [request responseString];
        NSLog(response);
        [[[UIAlertView alloc]initWithTitle:@"与你想Mate的用户列表" message:response delegate:nil cancelButtonTitle:@"取消" otherButtonTitles:nil, nil]show];
    }
    
}
@end
