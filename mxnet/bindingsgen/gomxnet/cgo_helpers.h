// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Wed, 07 Jun 2017 15:59:19 BST.
// By https://git.io/c-for-go. DO NOT EDIT.

#include "c_api.h"
#include "c_predict_api.h"
#include <stdlib.h>
#pragma once

#define __CGOGEN 1

// MXGenericCallback_2bc28855 is a proxy for callback MXGenericCallback.
int MXGenericCallback_2bc28855();

// CustomOpFBFunc_ff3016a2 is a proxy for callback CustomOpFBFunc.
int CustomOpFBFunc_ff3016a2(int arg0, void** arg1, int* arg2, int* arg3, int arg4, void* arg5);

// CustomOpDelFunc_7f9a5b2c is a proxy for callback CustomOpDelFunc.
int CustomOpDelFunc_7f9a5b2c(void* arg0);

// CustomOpListFunc_4e79662d is a proxy for callback CustomOpListFunc.
int CustomOpListFunc_4e79662d(char*** arg0, void* arg1);

// CustomOpInferShapeFunc_2ad17c00 is a proxy for callback CustomOpInferShapeFunc.
int CustomOpInferShapeFunc_2ad17c00(int arg0, int* arg1, unsigned int** arg2, void* arg3);

// CustomOpInferTypeFunc_59f65f2f is a proxy for callback CustomOpInferTypeFunc.
int CustomOpInferTypeFunc_59f65f2f(int arg0, int* arg1, void* arg2);

// CustomOpBwdDepFunc_7ac42004 is a proxy for callback CustomOpBwdDepFunc.
int CustomOpBwdDepFunc_7ac42004(int* arg0, int* arg1, int* arg2, int* arg3, int** arg4, void* arg5);

// CustomOpCreateFunc_ee34d11c is a proxy for callback CustomOpCreateFunc.
int CustomOpCreateFunc_ee34d11c(char* arg0, int arg1, unsigned int** arg2, int* arg3, int* arg4, struct MXCallbackList* arg5, void* arg6);

// CustomOpPropCreator_fdc4dbab is a proxy for callback CustomOpPropCreator.
int CustomOpPropCreator_fdc4dbab(char* arg0, int arg1, char** arg2, char** arg3, struct MXCallbackList* arg4);

// MXKVStoreUpdater_2e43854b is a proxy for callback MXKVStoreUpdater.
void MXKVStoreUpdater_2e43854b(int key, void* recv, void* local, void* handle);

// MXKVStoreServerController_98c9be52 is a proxy for callback MXKVStoreServerController.
void MXKVStoreServerController_98c9be52(int head, char* body, void* controller_handle);

