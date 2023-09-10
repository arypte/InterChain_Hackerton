pragma circom 2.1.4;


template UpdateWeight() {
    signal input x[2];          // 2차원 입력 벡터
    signal input y_actual;      // 실제 라벨
    signal input w_hidden[2][2];        // 숨겨진 노드의 2x2 가중치 행렬
    signal input w_output[2];           // 출력 노드의 2차원 가중치 벡터
    signal input alpha;                 // 학습률
    signal output w_hidden_updated[2][2]; // 업데이트 된 숨겨진 노드의 가중치
    signal output w_output_updated[2];     // 업데이트 된 출력 노드의 가중치

    // 숨겨진 레이어의 계산
    signal hidden[2];
    signal hidden_temp0 <== w_hidden[0][0] * x[0];
    signal hidden_temp1 <== w_hidden[0][1] * x[1];
    signal hidden_temp2 <== w_hidden[1][0] * x[0];
    signal hidden_temp3 <== w_hidden[1][1] * x[1];

    hidden[0] <== hidden_temp0 + hidden_temp1;
    hidden[1] <== hidden_temp2 + hidden_temp3;

    // 출력 계산: y_pred = w_output * hidden
    signal y_pred_temp0 <== w_output[0] * hidden[0];
    signal y_pred_temp1 <== w_output[1] * hidden[1];
    signal y_pred <== y_pred_temp0 + y_pred_temp1;

    // 오차: y_actual - y_pred
    signal error <== y_actual - y_pred;

    // 가중치에 대한 손실 함수의 편미분
    signal dL_dw_output[2];
    dL_dw_output[0] <== -hidden[0] * error;
    dL_dw_output[1] <== -hidden[1] * error;

    signal temp0, temp1, temp2, temp3;
    temp0 <== x[0] * error;
    temp1 <== x[1] * error;
    temp2 <== temp0 * w_output[0];
    temp3 <== temp1 * w_output[1];

    signal dL_dw_hidden[2][2];
    dL_dw_hidden[0][0] <== -temp2;
    dL_dw_hidden[0][1] <== -temp3;
    dL_dw_hidden[1][0] <== -temp0 * w_output[1];
    dL_dw_hidden[1][1] <== -temp1 * w_output[1];

    // 가중치 업데이트
    for (var i=0; i<2; i++) {
        w_output_updated[i] <== w_output[i] - alpha * dL_dw_output[i];
        for (var j=0; j<2; j++) {
            w_hidden_updated[i][j] <== w_hidden[i][j] - alpha * dL_dw_hidden[i][j];
        }
    }
}

component main{public [w_hidden,w_output,alpha]} = UpdateWeight();